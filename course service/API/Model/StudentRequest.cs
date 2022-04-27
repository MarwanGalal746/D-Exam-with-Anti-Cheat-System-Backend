using System.ComponentModel.DataAnnotations;

namespace API.Model;

public class StudentRequest
{
    [Required] public string StudentId { get; set; }
    [Required] public string CourseId { get; set; }
}