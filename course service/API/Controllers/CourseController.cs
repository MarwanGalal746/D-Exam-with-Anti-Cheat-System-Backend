using API.Model;
using Core.Entity;
using Core.Interface.Service;
using Core.Utility;
using Microsoft.AspNetCore.Mvc;

namespace API.Controllers;

[ApiController]
[Route("/api/courses/")]
public class CourseController : ControllerBase
{
    private readonly ICourseService _courseService;
    private readonly IStudCrsService _studCrsService;

    public CourseController(ICourseService courseService, IStudCrsService studCrsService)
    {
        _courseService = courseService;
        _studCrsService = studCrsService;
    }

    [HttpPost("/create-course")]
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

    [HttpGet("/get-course-by-id")]
    public async Task<IActionResult> GetById(string courseId)
    {
        return Ok(await _courseService.GetById(courseId));
    }
    
    [HttpGet("/get-course-by-name")]
    public async Task<IActionResult> GetByName(string name)
    {
        return Ok(await _courseService.GetByName(name));
    }
    
    [HttpGet("/get-course-by-code")]
    public async Task<IActionResult> GetByCode(string code)
    {
        return Ok(await _courseService.GetByCode(code));
    }
    
    [HttpGet("/get-teacher-courses")]
    public async Task<IActionResult> GetTeacherCourses(string teacherId)
    {
        var result = await _courseService.GetTeacherCourses(teacherId);
        Console.WriteLine("Started");
        foreach (var course in result)
        {
            Console.WriteLine("YES");
            Console.WriteLine(course);
        }
        return Ok();
    }
    
    [HttpPatch("/update-course")]
    public async Task<IActionResult> Update(UpdateCourseRequest request)
    {
        return Ok(await _courseService.Update(request.CourseId, request.NewCourseName));
    }
    
    [HttpDelete("/delete-course")]
    public IActionResult Delete(string id)
    {
        _courseService.Delete(id);
        return Ok("Success");
    }

    [HttpPost("/register-student")]
    public async Task<IActionResult> RegisterStudent(StudentRequest request)
    {
        await _studCrsService.AddStudentToCourse(request.StudentId, request.CourseId);
        await _courseService.RegisterStudent(request.StudentId, request.CourseId);
        return Ok("Success");
    }
    
    [HttpDelete("/remove-student")]
    public async Task<IActionResult> RemoveStudent(StudentRequest request)
    {
        await _studCrsService.RemoveStudentFromCourse(request.StudentId, request.CourseId);
        await _courseService.RemoveStudent(request.StudentId, request.CourseId);
        return Ok("Success");
    }
    
    [HttpGet("/get-student-courses")]
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
    
    [HttpPost("/create-student-courses")]
    public IActionResult CreateStudentCourses(string studentId)
    {
        _studCrsService.CreateStudentCourses(studentId);
        return Ok("Success");
    }
}