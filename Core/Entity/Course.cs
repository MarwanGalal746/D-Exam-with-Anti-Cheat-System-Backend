using System.ComponentModel.DataAnnotations;
using Core.Entity.Common;
using Redis.OM.Modeling;

namespace Core.Entity;

[Document(StorageType = StorageType.Json, Prefixes = new []{"Course"})]
public class Course : BaseEntity
{
    [RedisIdField] [Indexed] public Ulid CourseId { get; set; }
    [Indexed] [Required] public string TeacherId { get; set; }
    [Indexed] [Required] public string Name { get; set; } = null!;
    [Indexed] [Required] public string Code { get; set; } = null!;
    [Required] public string StartDate { get; set; }
    [Required] public string EndDate { get; set; }
    [Indexed] public List<string> Students { get; set; }
}