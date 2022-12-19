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

from pyspark.ml.feature import VectorAssembler

vecAssembler = VectorAssembler(inputCols=["total_reviews", "review_score"], outputCol="features")
#vecAssembler = VectorAssembler(inputCols=["id_genre", "review_score", "id_category"], outputCol="features")
data_r = data_r.drop("rewiew_id")
df_vec = vecAssembler.transform(data_r)
print(df_vec.count())


from pyspark.ml.clustering import KMeans
from pyspark.ml.evaluation import ClusteringEvaluator

kmeans = KMeans().setK(7).setSeed(1)
model = kmeans.fit(df_vec.select("features"))
# Make predictions
predictions = model.transform(df_vec)
df_t = model.transform(df_vec)
df_t.groupBy('prediction').count().show()
evaluator = ClusteringEvaluator()
silhouette = evaluator.evaluate(predictions)
print("Silhouette with squared euclidean distance = " + str(silhouette))


df_pd = df_t.toPandas()
f, (ax1) = plt.subplots(1, sharey=True,figsize=(8,8))
ax1.scatter(df_pd.total_reviews,df_pd.total_positive,c=df_pd.prediction,cmap='rainbow',  s=80, alpha=0.8)
plt.xlabel('Количество отзывов', fontsize=16)
plt.ylabel('Количество позитивных отзывов', fontsize=16)
plt.savefig('Clustering_rev.png')