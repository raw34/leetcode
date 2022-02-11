# 抽象工厂模式 Abstract factory

```
public interface IFactory {
    Mask create();
}


public class SurgicalMaskFactory implements IFactory{

    @Override
    public Mask create() {
        return new SurgicalMask();
    }
}


public class N95MaskFactory implements IFactory {
    @Override
    public Mask create() {
        return new N95Mask();
    }
}
```