# wifi

Displays WiFi signal level.

Available configuration keys:

| Key          | Description                   | Default value |
| ---          | ---                           | ---           |
| connected    | format for connected state    | {lvl}         |
| disconnected | format for disconnected state | N/A           |
| interface    | interface to monitor          | wlan0         |
| pause        | interval between updates      | 2000          |

Available variables for interpolation:

| Variable name | Description  |
| ---           | ---          |
| lvl           | signal level |

