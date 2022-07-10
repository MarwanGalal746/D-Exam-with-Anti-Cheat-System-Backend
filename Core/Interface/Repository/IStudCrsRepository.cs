using System.Linq.Expressions;
using Core.Entity;

namespace Core.Interface.Repository;

public interface IStudCrsRepository
{
    Task<StudentCourses?> GetStudentCourses(string studentId);
    Task<StudentCourses> Update(StudentCourses studentCourses);
    Task<StudentCourses> Create(StudentCourses studentCourses);
    Task<IList<StudentCourses>> FindBy(Expression<Func<StudentCourses, bool>> expression);
}