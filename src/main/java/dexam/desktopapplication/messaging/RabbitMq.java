package dexam.desktopapplication.messaging;

public class RabbitMq {
    private static final String QUEUE_NAME = "desktop-student";
    private static final String PORT= "5672";
    private static final String HOST = "20.25.175.236";
    public static boolean send(String message){
        try {
            String[] cmd = {"./RabbitMqPublisher.exe", "-host="+HOST, "-port="+PORT, "-queueName="+QUEUE_NAME, "-msg="+message};
            ProcessBuilder builder = new ProcessBuilder(cmd);
            builder.start();
        }catch (Exception e){
            System.out.println(e.getMessage());
            return false;
        }
        return true;
    }
}
