using System.ComponentModel.DataAnnotations;

namespace API.Model;

public class CreateCourseRequest
{
    [Required] public string Name { get; set; }
    [Required] public string TeacherId { get; set; }
    [Required] public DateTime StartDate { get; set; }
    [Required] public DateTime EndDate { get; set; }
}