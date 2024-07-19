package sql_test

import (
	"fmt"
	"github.com/jordanocokoljic/shrivel/v2/internal/sql"
	"testing"
)

func TestLexer_Next(t *testing.T) {
	tests := []struct {
		sql    string
		chunks []string
	}{
		{
			sql: "INSERT INTO customers (name, email) VALUES ('John Doe', 'john@example.com');",
			chunks: []string{
				"INSERT", "INTO", "customers", "(", "name", ",", "email", ")", "VALUES", "(", "'John Doe'", ",", "'john@example.com'", ")", ";",
			},
		},
		{
			sql: "UPDATE products \nSET price = price * 1.1, \n    last_updated = CURRENT_TIMESTAMP \nWHERE category = 'electronics' \n  AND stock_quantity < 10;",
			chunks: []string{
				"UPDATE", "products", "SET", "price", "=", "price", "*", "1.1", ",", "last_updated", "=", "CURRENT_TIMESTAMP", "WHERE", "category", "=", "'electronics'", "AND", "stock_quantity", "<", "10", ";",
			},
		},
		{
			sql: "WITH ranked_sales AS (\n    SELECT \n        product_id, \n        SUM(quantity) as total_quantity,\n        ROW_NUMBER() OVER (ORDER BY SUM(quantity) DESC) as rank\n    FROM sales\n    WHERE sale_date BETWEEN '2023-01-01' AND '2023-12-31'\n    GROUP BY product_id\n)\nSELECT \n    p.product_name,\n    r.total_quantity,\n    r.rank\nFROM ranked_sales r\nJOIN products p ON r.product_id = p.id\nWHERE r.rank <= 5\nORDER BY r.rank;",
			chunks: []string{
				"WITH", "ranked_sales", "AS", "(", "SELECT", "product_id", ",", "SUM", "(", "quantity", ")", "as", "total_quantity", ",", "ROW_NUMBER", "(", ")", "OVER", "(", "ORDER", "BY", "SUM", "(", "quantity", ")", "DESC", ")", "as", "rank", "FROM", "sales", "WHERE", "sale_date", "BETWEEN", "'2023-01-01'", "AND", "'2023-12-31'", "GROUP", "BY", "product_id", ")", "SELECT", "p", ".", "product_name", ",", "r", ".", "total_quantity", ",", "r", ".", "rank", "FROM", "ranked_sales", "r", "JOIN", "products", "p", "ON", "r", ".", "product_id", "=", "p", ".", "id", "WHERE", "r", ".", "rank", "<=", "5", "ORDER", "BY", "r", ".", "rank", ";",
			},
		},
		{
			sql: "CREATE TABLE orders\n(\n    order_id     INT PRIMARY KEY,\n    customer_id  INT,\n    order_date   DATE,\n    total_amount DECIMAL(10, 2)\n);",
			chunks: []string{
				"CREATE", "TABLE", "orders", "(", "order_id", "INT", "PRIMARY", "KEY", ",", "customer_id", "INT", ",", "order_date", "DATE", ",", "total_amount", "DECIMAL", "(", "10", ",", "2", ")", ")", ";",
			},
		},
		{
			sql: "SELECT first_name, last_name,\n       CASE\n           WHEN salary < 50000 THEN 'Low'\n           WHEN salary >= 50000 AND salary < 80000 THEN 'Medium'\n           ELSE 'High'\n       END AS salary_category\nFROM employees;\n",
			chunks: []string{
				"SELECT", "first_name", ",", "last_name", ",", "CASE", "WHEN", "salary", "<", "50000", "THEN", "'Low'", "WHEN", "salary", ">=", "50000", "AND", "salary", "<", "80000", "THEN", "'Medium'", "ELSE", "'High'", "END", "AS", "salary_category", "FROM", "employees", ";",
			},
		},
		{
			sql: "-- Get the names of all the citizens\nSELECT first, last FROM citizen;",
			chunks: []string{
				"SELECT", "first", ",", "last", "FROM", "citizen", ";",
			},
		},
		{
			sql: "/*\n The original query\n \n /*\n  Get all the first and last names\n  */\n SELECT first, last FROM citizen WHERE age >= 21;\n */\n \nSELECT first, last FROM citizen WHERE age >= 21;",
			chunks: []string{
				"SELECT", "first", ",", "last", "FROM", "citizen", "WHERE", "age", ">=", "21", ";",
			},
		},
		{
			sql: "SELECT title, content\nFROM articles\nWHERE MATCH(title, content) AGAINST ('database performance optimization' IN NATURAL LANGUAGE MODE);\n",
			chunks: []string{
				"SELECT", "title", ",", "content", "FROM", "articles", "WHERE", "MATCH", "(", "title", ",", "content", ")", "AGAINST", "(", "'database performance optimization'", "IN", "NATURAL", "LANGUAGE", "MODE", ")", ";",
			},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("Sample %d", i+1), func(t *testing.T) {
			lexer := sql.NewLexer([]rune(tt.sql))

			for i, chunk := range tt.chunks {
				read := lexer.Next()
				if string(read) != chunk {
					t.Errorf(
						"Unexpected chunk (%d). Expected: %q, got: %q",
						i, chunk, string(read))
				}
			}
		})
	}
}
