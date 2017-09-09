package com.dao;
import com.JDBC.SQLiteJDBC;
import java.sql.*;

public class abstruct_dao {

    protected static final String database_produce ="produce";

    protected static Connection conn = null;
    public static Connection getConn(){
        return conn;
    }

    public static void setConn(Connection conn){
        abstruct_dao.conn=conn;
    }

    public abstruct_dao(){
        try {
            if (conn==null) conn= SQLiteJDBC.getConnection();
            Statement stat = conn.createStatement();
            String sql = String.format("use %s;", database_produce);
            boolean flag=stat.execute(sql);
            if (flag) System.err.println("Can't find database \""+ database_produce +"\".");
        } catch (SQLException e) {
            e.printStackTrace();
            throw new RuntimeException(e);
        }
    }

    public abstruct_dao(Connection conn){
        try {
            conn = conn;
            Statement stat = conn.createStatement();
            String sql = String.format("use %s;", database_produce);
            boolean flag=stat.execute(sql);
            if (flag) System.err.println("Can't find database \""+ database_produce +"\".");
        } catch (SQLException e) {
            e.printStackTrace();
            throw new RuntimeException(e);
        }
    }

    public static void close(){
        /*
        * 内测阶段，啥都不干。
        * */
        //if (conn!=null) JDBC.close(conn,null);
        //conn=null;
    }

    public static void connect(){
        if (conn==null) conn= SQLiteJDBC.getConnection();
    }

    public static void work_begin(){
        /*
        * 用于开启一个事务
        * */
        connect();
        try {
            Statement stat = conn.createStatement();
            String sql = "BEGIN ;";
            boolean flag=stat.execute(sql);
            if (flag) System.err.println("Cannot begin a work.");
        } catch (SQLException e) {
            e.printStackTrace();
            throw new RuntimeException(e);
        }
    }

    public static void work_rollback(){
        /*
        * 用于回滚一个事务
        * */
        connect();
        try {
            Statement stat = conn.createStatement();
            String sql = "ROLLBACK ;";
            boolean flag=stat.execute(sql);
            if (flag) System.err.println("Cannot rollback a work.");
        } catch (SQLException e) {
            e.printStackTrace();
            throw new RuntimeException(e);
        }
    }

    public static void work_commit(){
        /*
        * 用于提交一个事务
        * */
        connect();
        try {
            Statement stat = conn.createStatement();
            String sql = "COMMIT ;";
            boolean flag=stat.execute(sql);
            if (flag) System.err.println("Cannot commit a work.");
        } catch (SQLException e) {
            e.printStackTrace();
            throw new RuntimeException(e);
        }
    }

    public static void runSQL(String sql){
        connect();
        try {
            Statement stat = conn.createStatement();
            boolean flag=stat.execute(sql);
            if (flag) System.err.println("Cannot run the sql "+sql+" .");
        } catch (SQLException e) {
            e.printStackTrace();
            throw new RuntimeException(e);
        }
    }

    public static void locktable(String tablename,boolean isReadLock){
        String lockname;
        if (isReadLock) lockname= "READ";
        else lockname="WRITE";
        String sql= String.format("LOCK TABLE %s %s", tablename, lockname);
        runSQL(sql);
        return ;
    }

    public static void unlock(String tablename){
        String sql= String.format("UNLOCK %s", tablename);
        runSQL(sql);
        return ;
    }

    public static void importFileLinux(String tablename,String filename){
        String sql=String.format("LOAD DATA INFILE '%s'\n" +
                "INTO TABLE %s\n" +
                "CHARACTER SET utf8\n" +
                "FIELDS TERMINATED BY ',' ENCLOSED BY '\"'\n" +
                "IGNORE 1 LINES;",filename,tablename);
        runSQL(sql);
    }

    public static void importFileWindows(String tablename,String filename){
        String sql=String.format("LOAD DATA INFILE \"%s\"\n" +
                "REPLACE INTO TABLE %s\n" +
                "CHARACTER SET gb2312\n" +
                "FIELDS TERMINATED BY \",\" ENCLOSED BY \"\"\n" +
                "LINES TERMINATED BY \"\\r\\n\"\n" +
                "IGNORE 1 LINES;",filename,tablename);
        runSQL(sql);
    }

}
