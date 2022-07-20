package hello;

import org.springframework.context.ApplicationContext;
import org.springframework.context.annotation.AnnotationConfigApplicationContext;
import org.springframework.context.annotation.ComponentScan;

@ComponentScan
public class ApplicationSpring {
    public static void main(String [] args){
        System.out.println("applicationSpring");
        //初始化Spring容器
        ApplicationContext context = new AnnotationConfigApplicationContext(ApplicationSpring.class);
        //从容器中获取对象
        MessagePrinter printer = context.getBean(MessagePrinter.class);

        System.out.println(printer);
        System.out.println(service);

        printer.printMessage();
    }
}
