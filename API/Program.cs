using Core.Interface.Repository;
using Core.Interface.Service;
using Persistence;
using Persistence.Config;
using Redis.OM;
using Service;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container.
builder.Services.AddControllers();

builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

builder.Services.AddSingleton(new RedisConnectionProvider(Environment.GetEnvironmentVariable("REDIS_CONNECTION_STRING", EnvironmentVariableTarget.Process)));
builder.Services.AddHostedService<IndexCreationService>();

builder.Services.AddSingleton<ICourseRepository, CourseRepository>();
builder.Services.AddSingleton<ICourseService, CourseService>();

builder.Services.AddSingleton<IStudCrsRepository, StudCrsRepository>();
builder.Services.AddSingleton<IStudCrsService, StudCrsService>();

builder.Services.AddCors();

var app = builder.Build();

//app.ConfigureExceptionHandler();

// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}

app.UseHttpsRedirection();

app.UseCors(c => c.AllowAnyHeader().AllowAnyMethod().AllowAnyOrigin());

app.UseAuthorization();

app.MapControllers();

app.Run();