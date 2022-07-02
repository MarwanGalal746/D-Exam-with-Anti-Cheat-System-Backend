using Core.Entity;
using Core.Interface.Repository;
using Core.Interface.Service;

namespace Service;

public class StudCrsService : IStudCrsService
{
    private readonly IStudCrsRepository _repository;

    public StudCrsService(IStudCrsRepository repository)
    {
        _repository = repository;
    }
    
    public void CreateStudentCourses(string studentId)
    {
        StudentCourses studentCourses = new StudentCourses
        {
            Courses = new List<string>(),
            StudentId = studentId,
            CreatedAt = DateTime.Now.ToString()
        };
        _repository.Create(studentCourses);
    }

    public async Task AddStudentToCourse(string studentId, string courseId)
    {
        var sc = await _repository.GetStudentCourses(studentId);
        if (sc != null)
        {
            sc.Courses.Add(courseId);
            await _repository.Update(sc);
        }
        else
        {
            throw new Exception("Cannot find student with the specified id: " + studentId);
        }
    }

    public async Task RemoveStudentFromCourse(string studentId, string courseId)
    {
        var sc = await _repository.GetStudentCourses(studentId);
        if (sc != null)
        {
            sc.Courses.Remove(courseId);
            await _repository.Update(sc);
        }
        else
        {
            throw new Exception("Cannot find student with the specified id: " + studentId);
        }
    }
    
    public async Task<List<string>> GetStudentCourses(string id)
    {
        var sc = await _repository.GetStudentCourses(id);
        return sc!.Courses;
    }
}