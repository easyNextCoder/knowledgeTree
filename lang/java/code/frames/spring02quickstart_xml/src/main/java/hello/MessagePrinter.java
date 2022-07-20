package hello;

import java.util.*;
/*
    打印机
 */
public class MessagePrinter {

    private  MessageService service;

    public MessagePrinter() {
        super();
        System.out.println("MessagePrinter.");
        List<Integer> con = new ArrayList<Integer>();

    }

    /*
        建立两个类的关联关系
     */

    public void setService(MessageService service) {
        this.service = service;
    }

    public void printMessage(){
        System.out.println(this.service.getMessage());
    }
}


