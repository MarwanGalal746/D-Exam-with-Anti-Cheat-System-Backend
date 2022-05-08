using System.ComponentModel.DataAnnotations;

namespace Core.Entity.Common;

public abstract class BaseEntity
{
    [Required] public string CreatedAt { get; set; }
    public string UpdatedAt { get; set; }
}