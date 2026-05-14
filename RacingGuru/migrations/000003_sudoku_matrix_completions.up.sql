CREATE TABLE public.sudoku_matrix_completions (
    user_id uuid NOT NULL,
    matrix_type text NOT NULL,
    completed_on date DEFAULT CURRENT_DATE NOT NULL
);

ALTER TABLE ONLY public.sudoku_matrix_completions
    ADD CONSTRAINT sudoku_matrix_completions_matrix_type_check
    CHECK (matrix_type IN ('drivers', 'teams'));

ALTER TABLE ONLY public.sudoku_matrix_completions
    ADD CONSTRAINT sudoku_matrix_completions_unique
    UNIQUE (user_id, matrix_type, completed_on);

ALTER TABLE ONLY public.sudoku_matrix_completions
    ADD CONSTRAINT sudoku_matrix_completions_user_id_fk
    FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;
