package com.dao.DBinterface;

import com.beans.produce;
import java.sql.SQLException;
import java.util.Date;
import java.util.List;

/**
 * Created by lee on 2017/9/1.
 *
 * @author: lee
 * create time: 下午1:51
 */
public interface ProduceInterface {
    // normal user and admin to get data
    // not crash to upper but pass empty

    public produce getProduceByID(int ID) throws SQLException;

    public List<produce> getProduceAccordRank(int Start, int end) throws SQLException;//排序

    public List<produce> searchProduce(String keyWord) throws SQLException;


}
