import { Component } from '@angular/core';
import { RouterLink } from '@angular/router';

@Component({
  selector: 'app-home-page',
  standalone: true,
  imports: [RouterLink],
  templateUrl: './home-page.component.html',
  styleUrl: './home-page.component.scss'
})
export class HomePageComponent {
  protected readonly carSlides = [
    {
      src: '/images/formula1.jpg',
      title: 'Formula pace',
      caption: 'Точные решения на скорости квалификационного круга.'
    },
    {
      src: '/images/wec.jpg',
      title: 'Endurance mind',
      caption: 'Длинная дистанция данных без потери концентрации.'
    },
    {
      src: '/images/gt3.jpeg',
      title: 'GT control',
      caption: 'Пилоты, команды и результаты в одном гоночном пульте.'
    },
    {
      src: '/images/lmp.jpg',
      title: 'Prototype data',
      caption: 'Быстрый доступ к статистике.'
    }
  ];

  protected readonly sections = [
    {
      title: 'Статистика без пит-стопов',
      description: 'Открывайте пилотов и команды, сравнивайте метрики и сразу видьте лучший показатель.'
    },
    {
      title: 'Личный гараж',
      description: 'Сохраняйте любимых гонщиков и команды, а затем возвращайтесь к ним в один клик.'
    },
    {
      title: 'Стартовая решётка',
      description: 'Проверь свои знания и займи поул!'
    }
  ];
}
