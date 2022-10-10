public class Test {
    public static void main(String[] args) throws IOException {
        byte[] KEY = new byte[]{//KEY
                0x00000000, 0x00000000,
                0x00000000, 0x00000000,
                0x00000000, 0x00000000,
                0x00000000, 0x00000000,
                0x00000000, 0x00000000,
                0x00000000, 0x00000000,
                0x00000000, 0x00000000,
                0x00000000, 0x00000000,
        };
        
        byte[] content = new byte[]{//加密内容
                0x00000000, 0x00000000,
                0x00000000, 0x00000000,
        };        
        
        TeaUtil teaUtil = new TeaUtil();
        byte[] enByte = teaUtil.encrypt(content,KEY); //加密后的字节
        byte[] deByte = teaUtil.decrypt(enByte,KEY); //解密后的字节
       
    }
}