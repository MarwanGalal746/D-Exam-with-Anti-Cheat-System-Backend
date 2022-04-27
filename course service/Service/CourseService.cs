using System.Diagnostics.CodeAnalysis;
using Core.Entity;
using Core.Interface.Repository;
using Core.Interface.Service;

namespace Service;

[SuppressMessage("ReSharper", "SpecifyACultureInStringConversionExplicitly")]
public class CourseService : ICourseService
{
    private readonly ICourseRepository _repository;

    public CourseService(ICourseRepository repository)
    {
        _repository = repository;
    }
    
    public async Task<Course> Create(Course course)
    {
        return await _repository.Create(course);
    }

    public async Task<Course> GetById(string id)
    {
        return (await _repository.Get(id))!;
    }

    public async Task<IList<Course>> GetByName(string name)
    {
        return await _repository.FindBy(c => c.Name == name);
    }

    public async Task<IList<Course>> GetByTeacher(string id)
    {
        return await _repository.FindBy(c => c.TeacherId == id);
    }

    public async Task<Course> GetByCode(string code)
    {
        var queryResult = await _repository.FindBy(c => c.Code == code);
        return queryResult.FirstOrDefault()!;
    }

    public async Task<Course> Update(string courseId, string newCourseName)
    {
        var course = await _repository.Get(courseId);
        if (course != null)
        {
            course.Name = newCourseName;
            course.UpdatedAt = DateTime.Now.ToString();
            return await _repository.Update(course);
        }
        throw new Exception("Cannot find course with the specified id: " + courseId);
    }

    public void Delete(string id)
    {
        _repository.Delete(id);
    }

    public async Task RegisterStudent(string studentId, string courseId)
    {
        var course = await _repository.Get(courseId);
        if (course != null)
        {
            course.Students.Add(studentId);
            await _repository.Update(course);
        }
        else
        {
            throw new Exception("Cannot find course with the specified id: " + courseId);
        }
    }
    
    public async Task RemoveStudent(string studentId, string courseId)
    {
        var course = await _repository.Get(courseId);
        if (course != null)
        {
            course.Students.Remove(studentId);
            await _repository.Update(course);
        }
        else
        {
            throw new Exception("Cannot find course with the specified id: " + courseId);
        }
    }
}