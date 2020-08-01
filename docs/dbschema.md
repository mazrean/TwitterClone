## twitter-clone

### users
| Name | Type | Null | Key | Default | Extra | 説明 |
| --- | --- | --- | --- | --- | --- | --- |
| id | int(11) | No | PRI |  | AUTO_INCREMENT |  |
| name | varchar(32) | No |  |  |  |  |
| password | char(128) | No |  |  |  |  |
| created_at | datetime(6) | No |  | CURRENT_TIMESTAMP |  |  |
| updated_at | datetime(6) | No |  | CURRENT_TIMESTAMP |  |  |
| deleted_at | datetime(6) |  |  |  |  |  |

### messages
| Name | Type | Null | Key | Default | Extra | 説明 |
| --- | --- | --- | --- | --- | --- | --- |
| id | varchar(36) | No | PRI |  |  |  |
| user_id | int(11) | No | MUL |  |  |  |
| content | text | No |  |  |  |  |
| is_pinned | boolean | No |  | false |  |  |
| created_at | datetime(6) | No |  | CURRENT_TIMESTAMP |  |  |
| updated_at | datetime(6) |  |  |  |  |  |
| deleted_at | datetime(6) |  |  |  |  |  |

### follows
| Name | Type | Null | Key | Default | Extra | 説明 |
| --- | --- | --- | --- | --- | --- | --- |
| id | int(11) | No | PRI |  | AUTO_INCREMENT |  |
| user_id | int(11) | No | MUL |  |  |  |
| target_user_id | int(11) | No | MUL |  |  |  |
| created_at | datetime(6) | No |  | CURRENT_TIMESTAMP |  |  |
| deleted_at | datetime(6) |  |  |  |  |  |

### favorites
| Name | Type | Null | Key | Default | Extra | 説明 |
| --- | --- | --- | --- | --- | --- | --- |
| id | int(11) | No | PRI |  | AUTO_INCREMENT |  |
| user_id | int(11) | No | MUL |  |  |  |
| target_message_id | varchar(36) | No |  |  |  |  |
| created_at | datetime(6) | No |  | CURRENT_TIMESTAMP |  |  |
| deleted_at | datetime(6) |  |  |  |  |  |
