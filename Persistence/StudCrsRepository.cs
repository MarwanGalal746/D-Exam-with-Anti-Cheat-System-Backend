using System.Linq.Expressions;
using Core.Entity;
using Core.Interface.Repository;
using Redis.OM;
using Redis.OM.Searching;

namespace Persistence;

public class StudCrsRepository : IStudCrsRepository
{
    private readonly RedisCollection<StudentCourses> _studentCourses;
    private readonly RedisConnectionProvider _provider;
    
    public StudCrsRepository(RedisConnectionProvider provider)
    {
        _provider = provider;
        _studentCourses =(RedisCollection<StudentCourses>)provider.RedisCollection<StudentCourses>();
    }
    public async Task<StudentCourses?> GetStudentCourses(string studentId)
    {
        return await _studentCourses.FindByIdAsync(studentId);
    }

    public async Task<StudentCourses> Update(StudentCourses studentCourses)
    {
        await _studentCourses.Update(studentCourses);
        await _studentCourses.SaveAsync();
        return studentCourses;
    }

    public async Task<StudentCourses> Create(StudentCourses studentCourses)
    {
        await _studentCourses.InsertAsync(studentCourses);
        return studentCourses;
    }
    
    public async Task<IList<StudentCourses>> FindBy(Expression<Func<StudentCourses, bool>> expression)
    {
        return await _studentCourses.Where(expression).ToListAsync();
    }
}