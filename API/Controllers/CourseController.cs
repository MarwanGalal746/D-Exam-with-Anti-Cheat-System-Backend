using System.Diagnostics.CodeAnalysis;
using API.Model;
using Core.Entity;
using Core.Interface.Service;
using Core.Utility;
using Messaging;
using Microsoft.AspNetCore.Mvc;

namespace API.Controllers;

[ApiController]
[Route("/api/courses/")]
[SuppressMessage("ReSharper", "SpecifyACultureInStringConversionExplicitly")]
public class CourseController : ControllerBase
{
    private readonly ICourseService _courseService;
    private readonly IStudCrsService _studCrsService;

    public CourseController(ICourseService courseService, IStudCrsService studCrsService)
    {
        _courseService = courseService;
        _studCrsService = studCrsService;
    }

    [HttpPost]
    [Route("create")]
    public async Task<IActionResult> Create(CreateCourseRequest request)
    {
        Course course = new Course
        {
            Name = request.Name,
            TeacherId = request.TeacherId,
            Code = CodeGenerator.Generate(),
            StartDate = request.StartDate.ToString(),
            EndDate = request.EndDate.ToString(),
            CreatedAt = DateTime.Now.ToString(),
            Students = new List<string>()
        };
        await _courseService.Create(course);
        
        return Ok(course);
    }

    [HttpGet]
    [Route("get-by-id")]
    public async Task<IActionResult> GetById(string courseId)
    {
        return Ok(await _courseService.GetById(courseId));
    }
    
    [HttpGet]
    [Route("get-by-name")]
    public async Task<IActionResult> GetByName(string name)
    {
        return Ok(await _courseService.GetByName(name));
    }
    
    [HttpGet]
    [Route("get-by-code")]
    public async Task<IActionResult> GetByCode(string code)
    {
        return Ok(await _courseService.GetByCode(code));
    }
    
    [HttpGet]
    [Route("get-by-teacher")]
    public async Task<IActionResult> GetByTeacher(string id)
    {
        return Ok(await _courseService.GetByTeacher(id));
    }
    
    [HttpPatch]
    [Route("update")]
    public async Task<IActionResult> Update(UpdateCourseRequest request)
    {
        return Ok(await _courseService.Update(request.CourseId, request.NewCourseName));
    }
    
    [HttpDelete]
    [Route("delete")]
    public IActionResult Delete(string courseId)
    {
        _courseService.Delete(courseId);
        RabbitMq.Send(courseId);
        return Ok("Success");
    }

    [HttpPost]
    [Route("register-student")]
    public async Task<IActionResult> RegisterStudent(RegisterStudentRequest request)
    {
        var course = await _courseService.GetByCode(request.CourseCode);
        await _studCrsService.AddStudentToCourse(request.StudentId, course.CourseId.ToString());
        await _courseService.RegisterStudent(request.StudentId, course.CourseId.ToString());
        return Ok("Success");
    }
    
    [HttpDelete]
    [Route("remove-student")]
    public async Task<IActionResult> RemoveStudent(RemoveStudentRequest request)
    {
        await _studCrsService.RemoveStudentFromCourse(request.StudentId, request.CourseId);
        await _courseService.RemoveStudent(request.StudentId, request.CourseId);
        return Ok("Success");
    }
    
    [HttpGet]
    [Route("get-student-courses")]
    public async Task<IActionResult> GetStudentCourses(string studentId)
    {
        var sc = await _studCrsService.GetStudentCourses(studentId);
        List<Course> courses = new List<Course>();
        
        foreach (var courseId in sc)
        {
            var course = await _courseService.GetById(courseId);
            courses.Add(course);
        }

        return Ok(courses);
    }
    
    [HttpPost]
    [Route("create-student-course")]
    public IActionResult CreateStudentCourses(string studentId)
    {
        _studCrsService.CreateStudentCourses(studentId);
        return Ok("Success");
    }
}