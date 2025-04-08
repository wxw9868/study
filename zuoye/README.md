### 1. AlertDialog的流程
AlertDialog是对话框中功能最强大、用法最灵活的一种，同时它也是其他3种对话框的父类。使用AlertDialog生成的对话框样式多变，但是基本样式总会包含4个区域：图标区、标题区、内容区、按钮区。  
创建一个AlertDialog一般需要如下几个步骤：  
（1）创建AlertDialog.Builder对象。  
（2）调用AlertDialog.Builder的setTitle()或setCustomTitle()方法设置标题。  
（3）调用AlertDialog.Builder的setIcon()设置图标。  
（4）调用AlertDialog.Builder的相关设置方法设置对话框内容。  
（5）调用AlertDialog.Builder的setPositiveButton()、setNegativeButton()或setNeutralButton()方法 添加多个按钮。  
（6）调用AlertDialog.Builder的create()方法创建AlertDialog对象，再调用AlertDialog对象的show()方法将该对话框显示出来。  

### 2. 自定义一个事件监听器
```java
import android.os.Bundle;
import android.view.View;
import android.widget.Button;
import android.widget.Toast;
import androidx.appcompat.app.AppCompatActivity;
public class MainActivity extends AppCompatActivity {
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        // 找到按钮视图
        Button button = findViewById(R.id.button);
        // 实现 OnClickListener 接口
        button.setOnClickListener(new View.OnClickListener()) {
            @Override
            public void onClick(View v) {
                // 事件处理逻辑：显示一个 Toast 消息
                Toast.makeText(MainActivity.this, "Button clicked", Toast.LENGTH_SHORT).show();
            }
        };
    }
}
```  
