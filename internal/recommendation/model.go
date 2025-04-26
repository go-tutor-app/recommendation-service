package recommendation

type StudentData struct {
	Code    string // Mã định danh của học sinh
	Content string // Nội dung liên quan (ví dụ: kết quả bài tập, hồ sơ, etc.)
}

func (u *StudentData) TableName() string {
	return "student_data"
}
