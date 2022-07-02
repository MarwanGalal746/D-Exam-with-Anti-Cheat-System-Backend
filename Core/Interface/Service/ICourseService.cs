using Core.Entity;

namespace Core.Interface.Service;

public interface ICourseService
{
    Task<Course> Create(Course course);
    Task<Course> GetById(string id);
    Task<IList<Course>> GetByName(string name);
    Task<IList<Course>> GetByTeacher(string id);
    Task<Course> GetByCode(string code);
    Task<Course> Update(string courseId, string newCourseName);
    void Delete(string id);
    Task RegisterStudent(string studentId, string courseId);
    Task RemoveStudent(string studentId, string courseId);
}