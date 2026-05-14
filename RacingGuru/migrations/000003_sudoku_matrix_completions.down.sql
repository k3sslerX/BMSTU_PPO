ALTER TABLE IF EXISTS ONLY public.sudoku_matrix_completions
    DROP CONSTRAINT IF EXISTS sudoku_matrix_completions_user_id_fk;

ALTER TABLE IF EXISTS ONLY public.sudoku_matrix_completions
    DROP CONSTRAINT IF EXISTS sudoku_matrix_completions_unique;

ALTER TABLE IF EXISTS ONLY public.sudoku_matrix_completions
    DROP CONSTRAINT IF EXISTS sudoku_matrix_completions_matrix_type_check;

DROP TABLE IF EXISTS public.sudoku_matrix_completions;
