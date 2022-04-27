using Core.Interface.Repository;
using Core.Interface.Service;
using Persistence;
using Persistence.Config;
using Redis.OM;
using Service;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container.
builder.Services.AddControllers();

// Learn more about configuring Swagger/OpenAPI at https://aka.ms/aspnetcore/swashbuckle
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

builder.Services.AddSingleton(new RedisConnectionProvider(builder.Configuration["REDIS_CONNECTION_STRING"]));
builder.Services.AddHostedService<IndexCreationService>();

builder.Services.AddSingleton<ICourseRepository, CourseRepository>();
builder.Services.AddSingleton<ICourseService, CourseService>();

builder.Services.AddSingleton<IStudCrsRepository, StudCrsRepository>();
builder.Services.AddSingleton<IStudCrsService, StudCrsService>();

builder.Services.AddCors();

var app = builder.Build();

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