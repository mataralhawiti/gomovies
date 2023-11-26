-- Github public dataset in BigQuery
-- standardSQL
SELECT pop, repo_name, path
FROM (
  SELECT id, repo_name, path
  FROM `bigquery-public-data.github_repos.files` AS files
  WHERE path LIKE '%pom.xml' AND
    EXISTS (
      SELECT 1
      FROM `bigquery-public-data.github_repos.contents`
      WHERE NOT binary AND
        content LIKE '%commons-collections<%' AND
        content LIKE '%>3.2.1<%' AND
        id = files.id
    )
)
JOIN (
  SELECT
    difference.new_sha1 AS id,
    ARRAY_LENGTH(repo_name) AS pop
  FROM `bigquery-public-data.github_repos.commits`
  CROSS JOIN UNNEST(difference) AS difference
)
USING (id)
ORDER BY pop DESC;

--
-- movies with highest voting
WITH Highest_voting As (
  SELECT movie_id, vote
  FROM `gomovies.moives_dtls`
  GROUP BY movie_id,vote
  HAVING vote = (SELECT MAX(vote) FROM `gomovies.moives_dtls`)
)
SELECT name, year, certificate, runtime, rating, `desc`, detls.vote
FROM `gomovies.moives_dtls` As detls, Highest_voting As v
WHERE detls.movie_id = v.movie_id;