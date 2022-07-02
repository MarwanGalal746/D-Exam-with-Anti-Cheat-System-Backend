using System.Linq.Expressions;
using Core.Entity;
using Core.Interface.Repository;
using Redis.OM;
using Redis.OM.Searching;

namespace Persistence;

public class CourseRepository : ICourseRepository
{
    private readonly RedisCollection<Course> _courses;
    private readonly RedisConnectionProvider _provider;
    
    public CourseRepository(RedisConnectionProvider provider)
    {
        _provider = provider;
        _courses = (RedisCollection<Course>)provider.RedisCollection<Course>();
    }
    public async Task<Course> Create(Course course)
    {
        await _courses.InsertAsync(course);
        return course;
    }

    public async Task<Course> Update(Course updatedCourse)
    {
        await _courses.Update(updatedCourse);
        await _courses.SaveAsync();
        return updatedCourse;
    }

    public void Delete(string id)
    {
        _provider.Connection.Unlink($"Course:{id}");
    }

    public async Task<Course?> Get(string id)
    {
        return await _courses.FindByIdAsync(id);
    }

    public async Task<IList<Course>> FindBy(Expression<Func<Course, bool>> expression)
    {
        return await _courses.Where(expression).ToListAsync();
    }
}