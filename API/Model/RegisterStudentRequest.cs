using System.ComponentModel.DataAnnotations;

namespace API.Model;

public class RegisterStudentRequest
{
    [Required] public string StudentId { get; set; }
    [Required] public string CourseCode { get; set; }
}