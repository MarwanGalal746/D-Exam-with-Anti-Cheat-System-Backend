using System.Text;
using Microsoft.AspNetCore.Authentication.JwtBearer;
using Microsoft.IdentityModel.Tokens;
using  Ocelot.DependencyInjection;
using Ocelot.Middleware;
 
var builder = WebApplication.CreateBuilder(args);
 
builder.Configuration.AddJsonFile(Environment.GetEnvironmentVariable("GATEWAY_JSON_FILE", EnvironmentVariableTarget.Process));

builder.Services.AddEndpointsApiExplorer();

builder.Services.AddCors();

builder.Services.AddAuthentication(options =>
{
    options.DefaultAuthenticateScheme = JwtBearerDefaults.AuthenticationScheme;
    options.DefaultChallengeScheme = JwtBearerDefaults.AuthenticationScheme;
}).AddJwtBearer(options =>
{
    options.RequireHttpsMetadata = false;
    options.SaveToken = true;
    options.TokenValidationParameters = new TokenValidationParameters
    {
        IssuerSigningKey = new SymmetricSecurityKey(Encoding.ASCII.GetBytes(Environment.GetEnvironmentVariable("AUTH_SECRET_KEY", EnvironmentVariableTarget.Process)!)),
        ValidateIssuerSigningKey = true,
        ValidateIssuer = false,
        ValidateAudience = false,
    };
});

builder.Services.AddOcelot();

var app = builder.Build();

app.UseHttpsRedirection();
app.UseAuthentication();
app.UseCors(c=>c.AllowAnyHeader().AllowAnyMethod().AllowAnyOrigin().WithExposedHeaders("Content-Disposition"));
app.UseOcelot().Wait();
app.UseAuthorization();

app.Run();