package rpn

import (
	"testing"
)

func TestCalculateRPN(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		want       float64
		wantErr    bool
	}{
		{
			name:       "Valid addition",
			expression: "3+4",
			want:       7,
			wantErr:    false,
		},
		{
			name:       "Valid subtraction",
			expression: "5-3",
			want:       2,
			wantErr:    false,
		},
		{
			name:       "Valid multiplication",
			expression: "2*3",
			want:       6,
			wantErr:    false,
		},
		{
			name:       "Valid division",
			expression: "12/4",
			want:       3,
			wantErr:    false,
		},
		{
			name:       "Division by zero",
			expression: "6/0",
			want:       0,
			wantErr:    true,
		},
		{
			name:       "Invalid expression",
			expression: "1+4-",
			want:       0,
			wantErr:    true,
		},
		{
			name:       "Unknown operator",
			expression: "3@4",
			want:       0,
			wantErr:    true,
		},
		{
			name:       "Start with operator",
			expression: "!2",
			want:       0,
			wantErr:    true,
		},
		{
			name:       "Have alpha",
			expression: "12+1a",
			want:       0,
			wantErr:    true,
		},
		{
			name:       "False input format",
			expression: "12)",
			want:       0,
			wantErr:    true,
		},
		{
			name:       "Сalculations with priority",
			expression: "12+12*(1+2)",
			want:       48,
			wantErr:    false,
		},
		{
			name:       "Сalculations with priority and unknown operator",
			expression: "12!+12*(1+2)",
			want:       0,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Calc(tt.expression)
			if (err != nil) != tt.wantErr {
				t.Errorf("Calc error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Calc = %v, want %v", got, tt.want)
			}
		})
	}
}
