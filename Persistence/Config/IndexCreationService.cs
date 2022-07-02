using Core.Entity;
using Microsoft.Extensions.Hosting;
using Redis.OM;

namespace Persistence.Config;

public class IndexCreationService : IHostedService
{
    private readonly RedisConnectionProvider _provider;
    public IndexCreationService(RedisConnectionProvider provider)
    {
        _provider = provider;
    }
    
    public async Task StartAsync(CancellationToken cancellationToken)
    {
        await _provider.Connection.CreateIndexAsync(typeof(Course));
        await _provider.Connection.CreateIndexAsync(typeof(StudentCourses));
    }

    public Task StopAsync(CancellationToken cancellationToken)
    {
        return Task.CompletedTask;
    }
}