# 工厂模式 Factory

## 简单工厂模式
```
public class MaskFactory {
    public Mask create(String type){
        // 使用简单工厂模式实现此处的逻辑
        switch (type){
            case "Surgical":
                return new SurgicalMask();
            case "N95":
                return new N95Mask();
            default:
                throw new IllegalArgumentException("Unsupported mask type");
        }
    }
}
```

## 工厂方法模式
```
public class SurgicalMaskFactory{

    public Mask create() {
        return new SurgicalMask();
    }
}


public class N95MaskFactory {
    public Mask create() {
        return new N95Mask();
    }
}
```