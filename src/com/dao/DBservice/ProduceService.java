package com.dao.DBservice;
import com.beans.produce;
import com.dao.DBinterface.ProduceInterface;
import com.dao.DBconnection.Dataconnect;
import java.sql.*;
import java.util.ArrayList;
import java.util.List;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class ProduceService implements ProduceInterface {

    public List<produce> finddata(ResultSet resultSet) throws Exception
    {
        List<produce> list= new ArrayList<>();
        while(resultSet.next()){
            produce produce=new produce();
            produce.setId(resultSet.getInt("id"));
            produce.setName(resultSet.getString("name"));
            produce.setAddition(resultSet.getString("addition"));
            produce.setLgroup(resultSet.getInt("lgroup"));
            produce.setLkind(resultSet.getInt("lkind"));
            produce.setLpercenta(resultSet.getInt("lpercenta"));
            produce.setLpercentb(resultSet.getInt("lpercentb"));
            produce.setLnumber(resultSet.getInt("lnumber"));
            produce.setMoney(resultSet.getDouble("money"));
            produce.setRank(resultSet.getInt("rank"));
            produce.setScore(resultSet.getDouble("score"));
            produce.setSgroup(resultSet.getInt("sgroup"));
            produce.setSkind(resultSet.getInt("skind"));
            produce.setSnumber(resultSet.getInt("snumber"));

            list.add(produce);
        }
        return list;
    }

    @Override
    public produce getProduceByID(int ID) throws SQLException {
        Dataconnect connect=new Dataconnect();
        ResultSet res = null;
        produce produce=new produce();
        PreparedStatement sta=null;
        List<produce> list=new ArrayList<>();
        try {
            Connection con=null;
            con=connect.getConnection();
            String sql="select * from Produce where id=?";
            sta=con.prepareStatement(sql);
            sta.setInt(1,ID);
            res=sta.executeQuery();
            list=finddata(res);
            System.out.println("查询成功");
        }
        catch (Exception e)
        {
            e.printStackTrace();
            System.out.println("数据库操作失败");
        }
        finally {
            sta.close();
        }
        if(list.isEmpty())
        {
            return null;
        }
        else return list.get(0);
    }


    @Override
    public List<produce> getProduceAccordRank(int Start, int end) throws SQLException {
        Dataconnect connect=new Dataconnect();
        ResultSet res = null;
        PreparedStatement sta=null;
        List<produce> list=new ArrayList<>();
        try {
            Connection con=null;
            con=connect.getConnection();
            System.out.print(con);
            String sql="select * from Produce ORDER BY rank  limit ?,?;";
            sta=con.prepareStatement(sql);
            sta.setInt(1,Start);
            sta.setInt(2,end-Start);
            System.out.println(sql);
            res=sta.executeQuery();
            list=finddata(res);
            System.out.println("查询成功");
        }
        catch (Exception e)
        {
            e.printStackTrace();
            System.out.println("数据库操作失败");
        }
        finally {
            sta.close();
        }
        return list;
    }

    @Override
    public List<produce> searchProduce(String keyWord) throws SQLException {
        Dataconnect connect=new Dataconnect();
        ResultSet res = null;
        PreparedStatement sta=null;
        List<produce> list=new ArrayList<>();
        try {
            Connection con=null;
            con=connect.getConnection();
            String sql="select * from Produce where name like '%"+keyWord+"%' ";
            sta=con.prepareStatement(sql);
            res=sta.executeQuery();
            list=finddata(res);
            System.out.println("查询成功");
        }
        catch (Exception e)
        {
            e.printStackTrace();
            System.out.println("数据库操作失败");
        }
        finally {
            sta.close();
        }
        return list;
    }




}
