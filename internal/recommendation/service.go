package recommendation

import (
	"context"
	"promotion/pkg/failure"
)

type StudentDataRepo struct {
	repo *Repo
}

func NewStudentDataRepo(repo *Repo) *StudentDataRepo {
	return &StudentDataRepo{repo: repo}
}

func (s *StudentDataRepo) GetStudentData(ctx context.Context, code string) (*StudentData, error) {
	data, err := s.repo.GetByCode(ctx, code)
	if err != nil {
		return nil, failure.ErrWithTrace(err)
	}
	return &StudentData{
		Code:    data.Code,
		Content: data.Content,
	}, nil
}

type ScoreCalculator struct {
	data *StudentDataRepo
}

func NewScoreCalculator(data *StudentDataRepo) *ScoreCalculator {
	return &ScoreCalculator{data: data}
}

func (c *ScoreCalculator) CalculateScore(ctx context.Context, code string) (int, error) {
	studentData, err := c.data.GetStudentData(ctx, code)
	if err != nil {
		return 0, err
	}
	score := len(studentData.Code)
	return score, nil
}

type LessonSuggester struct {
	data *StudentDataRepo
}

func NewLessonSuggester(data *StudentDataRepo) *LessonSuggester {
	return &LessonSuggester{data: data}
}

func (l *LessonSuggester) SuggestLessons(ctx context.Context, code string) ([]string, error) {
	studentData, err := l.data.GetStudentData(ctx, code)
	if err != nil {
		return nil, err
	}
	if len(studentData.Content) > 50 {
		return []string{"Advanced Lessons", "Challenge Tasks"}, nil
	}
	return []string{"Fundamentals", "Introduction"}, nil
}

type TrendAnalyzer struct {
	data *StudentDataRepo
}

func NewTrendAnalyzer(data *StudentDataRepo) *TrendAnalyzer {
	return &TrendAnalyzer{data: data}
}

func (t *TrendAnalyzer) AnalyzeTrends(ctx context.Context, code string) (string, error) {
	studentData, err := t.data.GetStudentData(ctx, code)
	if err != nil {
		return "", err
	}
	if len(studentData.Content)%2 == 0 {
		return "Positive Trend", nil
	}
	return "Needs Improvement", nil
}

// ------------------- Facade -------------------

type PerformanceFacade struct {
	scoreCalculator *ScoreCalculator
	lessonSuggester *LessonSuggester
	trendAnalyzer   *TrendAnalyzer
}

func NewPerformanceFacade(repo *Repo) *PerformanceFacade {
	studentDataRepo := NewStudentDataRepo(repo)
	return &PerformanceFacade{
		scoreCalculator: NewScoreCalculator(studentDataRepo),
		lessonSuggester: NewLessonSuggester(studentDataRepo),
		trendAnalyzer:   NewTrendAnalyzer(studentDataRepo),
	}
}

func (f *PerformanceFacade) GetStudentPerformance(ctx context.Context, code string) (int, string, error) {
	score, err := f.scoreCalculator.CalculateScore(ctx, code)
	if err != nil {
		return 0, "", err
	}
	trend, err := f.trendAnalyzer.AnalyzeTrends(ctx, code)
	if err != nil {
		return 0, "", err
	}
	return score, trend, nil
}

func (f *PerformanceFacade) GetLessonRecommendations(ctx context.Context, code string) ([]string, error) {
	return f.lessonSuggester.SuggestLessons(ctx, code)
}
