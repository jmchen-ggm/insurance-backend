package test;
import com.dao.DBinterface.ProduceInterface;
import com.dao.DBservice.ProduceService;
import com.beans.produce;
import java.util.List;


public class daotest {
    public static void main(String args[]) throws Exception {
        ProduceInterface Daouser = new ProduceService();
        //List<String> x = Daouser.getWORDAccordToWORDOffset(1, 500);
        String i = "1";
        List<produce> a = Daouser.getProduceAccordRank(0,10);
        //List<produce> a = Daouser.searchProduce("B");
        //System.out.println(a.getName());
        int y = 0;
        for (produce u : a) {
            y++;
            System.out.println(u.getName() + "  " + u.getRank());


        }

    }
}
