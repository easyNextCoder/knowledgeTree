package hello;

import org.springframework.stereotype.Component;

/*
    打印服务


 */
@Component
public class MessageService {
    public MessageService() {
        super();
        System.out.print("MessageService.");
    }

    public String getMessage(){
        return "hello world!";
    }
}
