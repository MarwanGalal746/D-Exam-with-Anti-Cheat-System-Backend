using Core.Entity;

namespace Core.Interface.Service;

public interface IStudCrsService
{
    void CreateStudentCourses(string studentId);
    public Task AddStudentToCourse(string studentId, string courseId);
    Task<List<string>> GetStudentCourses(string id);
    Task RemoveStudentFromCourse(string studentId, string courseId);
    Task RemoveStudentsFromSpecificCourse(string courseId);
}