import psycopg2
import pyspark
from pyspark.sql import SparkSession
from pyspark.sql import DataFrameReader
from pyspark.ml.feature import VectorAssembler
from pyspark.sql.types import *
from matplotlib import pyplot as plt
import seaborn as sns

con = psycopg2.connect(
  database="GameDevBackup", 
  user="postgres", 
  password="1111", 
  host="127.0.0.1", 
  port="5432"
)

print("Database opened successfully")

cur = con.cursor()  

sql = "COPY (SELECT * FROM \"DB_schema\".\"Reviews\") TO STDOUT WITH CSV DELIMITER ';'"
with open("table_Reviews3.csv", "w", encoding="utf-8") as file:
  cur.copy_expert(sql, file)


spark = SparkSession.builder\
        .master("local[*]")\
        .appName('PySpark')\
        .getOrCreate()



# добавляем таблицу оценок

import pyspark.sql.functions as f
from pyspark.sql.functions import col, when

data_schema_r = [
               StructField('appid', IntegerType(), True),
               StructField('total_reviews', IntegerType(), True),
               StructField('total_positive', IntegerType(), True),
               StructField('total_negative', IntegerType(), True),
               StructField('review_score', IntegerType(), True),
               StructField('review_score_desc', StringType(), True),
               StructField('rewiew_id', IntegerType(), True),
            ]

struc_r = StructType(fields = data_schema_r)

data_r = spark.read.csv(
    'table_Reviews3.csv',
    sep=';',
    header=True,
    schema=struc_r
)

# очищаем таблицу оценок, чтобы оценки были
data_r = data_r.filter(col("review_score") != 0)
print(data_r.count())
data_r.show(5)


data_r = data_r.filter(col("total_reviews") <= 25000)



vecAssembler = VectorAssembler(inputCols=["review_score", "total_reviews"], outputCol="features")
data_r = data_r.drop("rewiew_id")
df_vec = vecAssembler.transform(data_r)
print(df_vec.count())

df_pd = data_r.toPandas()
ax = sns.heatmap(df_pd.corr(), annot=True)
plt.savefig('Correlation.png')