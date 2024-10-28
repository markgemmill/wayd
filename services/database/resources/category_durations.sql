select 
  categories.name as name, 
  COUNT(entries.id) as entry_count,
  SUM(entries.duration) as seconds,
  format('%02u:', (floor(SUM(entries.duration / 3600.0)))) ||
  format('%02u:', (floor(SUM(entries.duration) / 60.0) - (floor(SUM(entries.duration / 3600.0)) * 60))) ||
  format('%02u', (SUM(entries.duration) - (floor(SUM(entries.duration) / 60.0) * 60)))
  as duration 
FROM entries 
left join projects on 
  projects.id = entries.project_id 
left join categories on
  projects.category_id = categories.id
WHERE 
  entries.duration > 0 AND 
  entries.start > ? AND 
  entries.end < ? 
GROUP BY categories.name
