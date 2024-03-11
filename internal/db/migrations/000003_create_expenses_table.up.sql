CREATE TABLE IF NOT EXISTS expenses (
    id UUID PRIMARY KEY,
    name VARCHAR(255),
    description TEXT,
    amount DOUBLE PRECISION DEFAULT 0,
    budget_id UUID NOT NULL,
    user_id UUID NOT NULL,
    CONSTRAINT fk_expense_user_id FOREIGN KEY(user_id) REFERENCES users(id),
    CONSTRAINT fk_expense_budget_id FOREIGN KEY(budget_id) REFERENCES budgets(id)
)