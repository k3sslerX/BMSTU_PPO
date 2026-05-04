import { Component, ElementRef, HostListener, Input, OnChanges, OnDestroy, SimpleChanges, forwardRef, inject, signal } from '@angular/core';
import { ControlValueAccessor, NG_VALUE_ACCESSOR } from '@angular/forms';

export interface LookupOption {
  id: string;
  label: string;
  subtitle?: string;
  searchText?: string;
}

@Component({
  selector: 'app-admin-lookup-combo',
  standalone: true,
  providers: [
    {
      provide: NG_VALUE_ACCESSOR,
      useExisting: forwardRef(() => AdminLookupComboComponent),
      multi: true
    }
  ],
  template: `
    <div class="combo-field" [class.is-disabled]="disabled()" [class.is-multiple]="multiple">
      @if (label) {
        <span class="field-label">{{ label }}</span>
      }

      <div class="combo-control">
        <input
          type="text"
          [value]="query()"
          [placeholder]="placeholder"
          [disabled]="disabled()"
          autocomplete="off"
          (focus)="openDropdown()"
          (input)="handleInput($any($event.target).value)"
        />
      </div>

      @if (multiple && selectedOptions().length) {
        <div class="selected-chips">
          @for (option of selectedOptions(); track option.id) {
            <button class="selected-chip" type="button" [disabled]="disabled()" (click)="removeOption(option.id)">
              <span>{{ option.label }}</span>
              <strong>×</strong>
            </button>
          }
        </div>
      }
    </div>
  `,
  styles: [
    `
      :host {
        display: block;
        position: relative;
        z-index: 1;
        min-width: 0;
      }

      :host:focus-within {
        z-index: 100000;
      }

      .combo-field {
        position: relative;
        z-index: 1;
        display: flex;
        flex-direction: column;
        gap: 0.45rem;
      }

      .combo-field:focus-within {
        z-index: 100000;
      }

      .field-label {
        font-size: 0.92rem;
        color: var(--muted);
      }

      .combo-control {
        display: block;
      }

      .combo-empty {
        color: var(--muted);
      }

      .combo-empty {
        padding: 0.8rem;
      }

      .selected-chips {
        display: flex;
        flex-wrap: wrap;
        gap: 0.45rem;
      }

      .selected-chip {
        display: inline-flex;
        align-items: center;
        gap: 0.4rem;
        max-width: 100%;
        border: 1px solid rgba(229, 9, 20, 0.2);
        border-radius: 999px;
        background: rgba(229, 9, 20, 0.1);
        color: var(--brand-strong);
        padding: 0.35rem 0.55rem 0.35rem 0.7rem;
      }

      .selected-chip span {
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
      }

      .selected-chip strong {
        color: var(--danger);
      }

      .is-disabled {
        opacity: 0.72;
      }
    `
  ]
})
export class AdminLookupComboComponent implements ControlValueAccessor, OnChanges, OnDestroy {
  private readonly host = inject(ElementRef<HTMLElement>);

  @Input() label = '';
  @Input() placeholder = 'Начните вводить название или UUID';
  @Input() loadingText = 'Загружаем список...';
  @Input() emptyText = 'Совпадений нет';
  @Input() allSelectedText = 'Все варианты уже выбраны';
  @Input() selectedPrefix = 'Выбрано';
  @Input() loading = false;
  @Input() multiple = false;
  @Input() closeOnSelect = false;
  @Input() options: LookupOption[] = [];

  protected readonly query = signal('');
  protected readonly dropdownOpen = signal(false);
  protected readonly disabled = signal(false);
  protected readonly selectedIds = signal<string[]>([]);

  private onChange: (value: string | string[]) => void = () => undefined;
  private onTouched: () => void = () => undefined;
  private overlay: HTMLElement | null = null;

  ngOnChanges(changes: SimpleChanges): void {
    if (changes['options'] && !this.multiple) {
      this.syncQueryFromSelected();
    }
    this.renderOverlay();
  }

  ngOnDestroy(): void {
    this.removeOverlay();
  }

  writeValue(value: unknown): void {
    const ids = Array.isArray(value)
      ? value.map((item) => normalizeSourceValue(item)).filter(Boolean)
      : normalizeSourceValue(value)
        ? [normalizeSourceValue(value)]
        : [];
    this.selectedIds.set(ids);
    if (!this.multiple) {
      this.syncQueryFromSelected();
    }
    this.renderOverlay();
  }

  registerOnChange(fn: (value: string | string[]) => void): void {
    this.onChange = fn;
  }

  registerOnTouched(fn: () => void): void {
    this.onTouched = fn;
  }

  setDisabledState(isDisabled: boolean): void {
    this.disabled.set(isDisabled);
    if (isDisabled) {
      this.closeDropdown();
    }
  }

  @HostListener('document:pointerdown', ['$event'])
  protected handleDocumentPointerDown(event: PointerEvent): void {
    const target = event.target as Node | null;
    if (!target || this.host.nativeElement.contains(target) || this.overlay?.contains(target)) {
      return;
    }

    this.closeDropdown();
  }

  @HostListener('window:resize')
  @HostListener('window:scroll')
  protected repositionDropdown(): void {
    this.renderOverlay();
  }

  protected openDropdown(): void {
    if (this.disabled()) {
      return;
    }
    this.dropdownOpen.set(true);
    this.onTouched();
    this.renderOverlay();
  }

  protected handleInput(value: string): void {
    this.query.set(value);
    this.dropdownOpen.set(true);
    this.onTouched();

    if (this.multiple) {
      this.renderOverlay();
      return;
    }

    const match = this.findExactOption(value);
    this.selectedIds.set(match ? [match.id] : []);
    this.onChange(match?.id ?? '');
    this.renderOverlay();
  }

  protected selectOption(option: LookupOption): void {
    if (this.multiple) {
      this.selectedIds.update((ids) => (ids.includes(option.id) ? ids : [...ids, option.id]));
      this.query.set('');
      this.emitValue();
      this.onTouched();
      if (this.closeOnSelect) {
        this.closeDropdown();
        return;
      }
      this.dropdownOpen.set(true);
      this.renderOverlay();
      return;
    }

    this.selectedIds.set([option.id]);
    this.query.set(option.label);
    this.emitValue();
    this.onTouched();
    this.closeDropdown();
  }

  protected removeOption(id: string): void {
    this.selectedIds.update((ids) => ids.filter((item) => item !== id));
    this.emitValue();
    this.onTouched();
    this.renderOverlay();
  }

  protected filteredOptions(): LookupOption[] {
    const selected = new Set(this.selectedIds());
    const normalizedQuery = normalize(this.query());

    return this.options.filter((option) => {
      if (this.multiple && selected.has(option.id)) {
        return false;
      }

      if (!normalizedQuery) {
        return true;
      }

      return normalize(this.optionSearchText(option)).includes(normalizedQuery);
    });
  }

  protected selectedOption(): LookupOption | null {
    const id = this.selectedIds()[0];
    return id ? this.findOptionById(id) ?? fallbackOption(id) : null;
  }

  protected selectedOptions(): LookupOption[] {
    return this.selectedIds().map((id) => this.findOptionById(id) ?? fallbackOption(id));
  }

  private emitValue(): void {
    this.onChange(this.multiple ? this.selectedIds() : this.selectedIds()[0] ?? '');
  }

  private syncQueryFromSelected(): void {
    const selected = this.selectedOption();
    this.query.set(selected?.label ?? this.selectedIds()[0] ?? '');
  }

  private findOptionById(id: string): LookupOption | undefined {
    return this.options.find((option) => option.id === id);
  }

  private findExactOption(value: string): LookupOption | undefined {
    const normalizedValue = normalize(value);
    if (!normalizedValue) {
      return undefined;
    }

    const idMatch = this.options.find((option) => normalize(option.id) === normalizedValue);
    if (idMatch) {
      return idMatch;
    }

    const labelMatches = this.options.filter((option) => normalize(option.label) === normalizedValue);
    return labelMatches.length === 1 ? labelMatches[0] : undefined;
  }

  private optionSearchText(option: LookupOption): string {
    return [option.label, option.subtitle, option.id, option.searchText].filter(Boolean).join(' ');
  }

  private closeDropdown(): void {
    this.dropdownOpen.set(false);
    this.removeOverlay();
  }

  private renderOverlay(): void {
    if (!this.dropdownOpen() || this.disabled()) {
      this.removeOverlay();
      return;
    }

    const overlay = this.ensureOverlay();
    this.positionOverlay(overlay);
    overlay.replaceChildren();

    if (this.loading) {
      overlay.appendChild(this.createEmptyItem(this.loadingText));
      return;
    }

    const options = this.filteredOptions();
    if (!options.length) {
      overlay.appendChild(
        this.createEmptyItem(this.multiple && this.options.length > 0 && this.selectedIds().length === this.options.length ? this.allSelectedText : this.emptyText)
      );
      return;
    }

    for (const option of options) {
      overlay.appendChild(this.createOptionButton(option));
    }
  }

  private ensureOverlay(): HTMLElement {
    if (this.overlay) {
      return this.overlay;
    }

    const overlay = document.createElement('div');
    overlay.className = 'admin-combo-overlay';
    document.body.appendChild(overlay);
    this.overlay = overlay;
    return overlay;
  }

  private positionOverlay(overlay: HTMLElement): void {
    const control = this.host.nativeElement.querySelector('.combo-control') ?? this.host.nativeElement;
    const rect = control.getBoundingClientRect();
    overlay.style.left = `${rect.left}px`;
    overlay.style.top = `${rect.bottom + 7}px`;
    overlay.style.width = `${rect.width}px`;
  }

  private createOptionButton(option: LookupOption): HTMLButtonElement {
    const button = document.createElement('button');
    button.type = 'button';
    button.className = 'admin-combo-option';
    button.addEventListener('click', () => this.selectOption(option));

    const title = document.createElement('strong');
    title.textContent = option.label;
    button.appendChild(title);

    if (option.subtitle) {
      const subtitle = document.createElement('span');
      subtitle.textContent = option.subtitle;
      button.appendChild(subtitle);
    }

    return button;
  }

  private createEmptyItem(text: string): HTMLSpanElement {
    const item = document.createElement('span');
    item.className = 'admin-combo-empty';
    item.textContent = text;
    return item;
  }

  private removeOverlay(): void {
    this.overlay?.remove();
    this.overlay = null;
  }
}

function normalize(value: string): string {
  return normalizeSourceValue(value).toLowerCase();
}

function normalizeSourceValue(value: unknown): string {
  return String(value ?? '').trim();
}

function fallbackOption(id: string): LookupOption {
  return {
    id,
    label: id
  };
}
