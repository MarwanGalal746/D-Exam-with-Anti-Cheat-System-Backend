using System.Text;

namespace Core.Utility;

public static class CodeGenerator
{
    private static readonly Random _random = new Random();
    private static int RandomNumber(int min, int max)  
    {  
        return _random.Next(min, max);  
    }
    private static string RandomString(int size, bool lowerCase = false)  
    {  
        var builder = new StringBuilder(size);
        char offset = lowerCase ? 'a' : 'A';  
        const int lettersOffset = 26; // A...Z or a..z: length = 26  
  
        for (var i = 0; i < size; i++)  
        {  
            var @char = (char)_random.Next(offset, offset + lettersOffset);  
            builder.Append(@char);  
        }  
  
        return lowerCase ? builder.ToString().ToLower() : builder.ToString();  
    }
    public static string Generate()  
    {  
        var code = new StringBuilder();  
  
        // 4-Letters lower case   
        code.Append(RandomString(2));  
  
        // 4-Digits between 1000 and 9999  
        code.Append(RandomNumber(10, 99));  
  
        // 2-Letters upper case  
        code.Append(RandomString(2));  
        
        return code.ToString();  
    }  
}  
