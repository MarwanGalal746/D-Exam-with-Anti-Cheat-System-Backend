using Core.Entity.Common;
using Redis.OM.Modeling;

namespace Core.Entity;

[Document(StorageType = StorageType.Json, Prefixes = new []{"StudentCourses"})]
public class StudentCourses : BaseEntity
{
    [RedisIdField] [Indexed] public string StudentId { get; set; }
    [Indexed] public List<string> Courses { get; set; } = null!;
}