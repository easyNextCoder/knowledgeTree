package hello;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

/*
    打印机
 */
@Component
public class MessagePrinter {


    public MessagePrinter() {
        super();
        System.out.println("MessagePrinter.");
    }

    private  MessageService servce;
    /*
        建立两个类的关联关系
     */
    @Autowired
    public void setServce(MessageService servce) {
        this.servce = servce;
    }

    public void printMessage(){
        System.out.println(this.servce.getMessage());
    }
}
