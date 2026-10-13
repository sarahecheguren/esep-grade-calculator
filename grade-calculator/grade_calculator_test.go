package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeC(t *testing.T) {
        expected_value := "C"
        
        gradeCalculator := NewGradeCalculator()

        gradeCalculator.AddGrade("open source assignment", 70, Assignment)
        gradeCalculator.AddGrade("exam 1", 71, Exam)
        gradeCalculator.AddGrade("essay on ai ethics", 77, Essay)
                
        actual_value := gradeCalculator.GetFinalGrade()
 
        if expected_value != actual_value {
                t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
        }
}

func TestGetGradeD(t *testing.T) {
        expected_value := "D"

        gradeCalculator := NewGradeCalculator()

        gradeCalculator.AddGrade("open source assignment", 62, Assignment)
        gradeCalculator.AddGrade("exam 1", 63, Exam)
        gradeCalculator.AddGrade("essay on ai ethics", 66, Essay)

        actual_value := gradeCalculator.GetFinalGrade()

        if expected_value != actual_value {
                t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value) 
        }
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 50, Assignment)
	gradeCalculator.AddGrade("exam 1", 45, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 51, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGradeType(t *testing.T) {
	cases := map[GradeType]string{
		Assignment: "assignment",
		Exam: "exam",
		Essay: "essay",
	}

	for gradeType, expected := range cases {
		if actual := gradeType.String(); actual != expected {
		t.Errorf("Expected String to return '%s'; got '%s' instead", expected, actual)
		}
	}
}
