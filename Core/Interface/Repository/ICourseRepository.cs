using System.Linq.Expressions;
using Core.Entity;

namespace Core.Interface.Repository;

public interface ICourseRepository
{
    Task<Course> Create(Course course);
    Task<Course> Update(Course updatedCourse);
    void Delete(string id);
    Task<Course?> Get(string id);
    Task<IList<Course>> FindBy(Expression<Func<Course, bool>> expression);
}