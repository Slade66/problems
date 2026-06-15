# 20260604174731. 互联链路本端对端配对

- **难度：** 🟡 中等

- **题型：** SQL · self-join

## 题目描述

> 遇到该问题的真实业务场景，以及从中抽象出来的纯净算法问题。

### 问题背景

- **互联链路是什么：**

    - 两台交换机之间连一根物理线，叫一条互联链路，线的两头各有一个端口和一个 IP。

    - **示意图：**

        ```
        ┌──────────────────┐                              ┌──────────────────┐
        │      S6_B        │                              │      LF6_A       │
        │   100GE1/1/1     │────────── 互联链路 ───────────│   100GE1/0/1     │
        │  172.16.1.1/30   │                              │  172.16.1.2/30   │
        └──────────────────┘                              └──────────────────┘
                本端                                              对端
        ```

- **系统如何存储：** 

    - 规划系统把一条链路拆成两张表，`connections` 记物理连线，`interconnects` 每端各存一行，只记本端、不存对端。

    - **`connections`（一根线一行）：**

        | connection_uuid | from_device | from_port  | to_device | to_port    |
        |-----------------|-------------|------------|-----------|------------|
        | conn-SA         | S6_B        | 100GE1/1/1 | LF6_A     | 100GE1/0/1 |

    - **`interconnects`（一端一行）：**

        | uuid    | connection_uuid | device | ip_address    |
        |---------|-----------------|--------|---------------|
        | ic-SA-s | conn-SA         | S6_B   | 172.16.1.1/30 |
        | ic-SA-a | conn-SA         | LF6_A  | 172.16.1.2/30 |

- **导致的问题：**

    - 数据库以每端一行规范化存储，对端信息不存；前端每一行却要同时显示本端与对端——两者之间存在结构差异，必须先把同一根线的两行配成一对返回给前端。

### 核心问题

- **输入：**

    - **`connections`（一根线一行）：**

        | connection_uuid | from_device | from_port  | to_device | to_port    |
        |-----------------|-------------|------------|-----------|------------|
        | conn-SA         | S6_B        | 100GE1/1/1 | LF6_A     | 100GE1/0/1 |

    - **`interconnects`（一端一行）：**

        | uuid    | connection_uuid | device | ip_address    |
        |---------|-----------------|--------|---------------|
        | ic-SA-s | conn-SA         | S6_B   | 172.16.1.1/30 |
        | ic-SA-a | conn-SA         | LF6_A  | 172.16.1.2/30 |

- **输出：** `interconnects` 每行一条结果

    | 列            | 含义     | 来源                                         |
    |---------------|----------|----------------------------------------------|
    | `ic_uuid`     | 本端 uuid | `interconnects.uuid`                        |
    | `local_device`| 本端设备 | `interconnects.device`                       |
    | `local_port`  | 本端接口 | 本端是 from → `from_port`，否则 `to_port`    |
    | `local_ip`    | 本端 IP  | `interconnects.ip_address`                   |
    | `peer_device` | 对端设备 | 本端是 from → `to_device`，否则 `from_device`|
    | `peer_port`   | 对端接口 | 本端是 from → `to_port`，否则 `from_port`    |
    | `peer_ip`     | 对端 IP  | 配对行的 `ip_address`；配不到则 `NULL`        |

- **配对规则：**

    - 对端行 = 同一个 `connection_uuid`、且 `device` 与本端不同的那行。

    - 两台设备之间可能拉了多根线，只用设备名配对会乱配——`connection_uuid` 是每根线的唯一标识，必须用它把各根线隔开。

## 示例解析

### 示例 1：单链路基本配对

- **输入：**

    - **`connections`：**

        | connection_uuid | from_device | from_port  | to_device | to_port    |
        |-----------------|-------------|------------|-----------|------------|
        | conn-1          | S6_B        | 100GE1/1/1 | LF6_A     | 100GE1/0/1 |

    - **`interconnects`：**

        | uuid | connection_uuid | device | ip_address    |
        |------|-----------------|--------|---------------|
        | ic-1 | conn-1          | S6_B   | 172.16.1.1/30 |
        | ic-2 | conn-1          | LF6_A  | 172.16.1.2/30 |

- **输出：**

    | ic_uuid | local_device | local_port | local_ip      | peer_device | peer_port  | peer_ip       |
    |---------|--------------|------------|---------------|-------------|------------|---------------|
    | ic-1    | S6_B         | 100GE1/1/1 | 172.16.1.1/30 | LF6_A       | 100GE1/0/1 | 172.16.1.2/30 |
    | ic-2    | LF6_A        | 100GE1/0/1 | 172.16.1.2/30 | S6_B        | 100GE1/1/1 | 172.16.1.1/30 |

- **解释：**

    - **ic-1（S6_B）：** 在 connections 里 S6_B 是 from_device，所以本端口取 `from_port` = 100GE1/1/1，对端设备取 `to_device` = LF6_A，对端接口取 `to_port` = 100GE1/0/1；再去 interconnects 里找同一 conn-1 下设备不是 S6_B 的那行，找到 ic-2，取其 ip_address = 172.16.1.2/30 作为 peer_ip

    - **ic-2（LF6_A）：** LF6_A 是 to_device，方向反过来，本端口取 `to_port`，对端取 from 那侧；peer_ip 来自 ic-1 的 ip_address = 172.16.1.1/30

### 示例 2：并行链路——同设备对之间有两根线

- **输入：**

    - **`connections`：**

        | connection_uuid | from_device | from_port  | to_device | to_port    |
        |-----------------|-------------|------------|-----------|------------|
        | conn-A          | LF6_A       | 100GE1/0/2 | LF6_C     | 100GE1/0/2 |
        | conn-B          | LF6_A       | 100GE1/0/3 | LF6_C     | 100GE1/0/3 |

    - **`interconnects`：**

        | uuid  | connection_uuid | device | ip_address     |
        |-------|-----------------|--------|----------------|
        | ic-A1 | conn-A          | LF6_A  | 172.16.10.1/30 |
        | ic-C1 | conn-A          | LF6_C  | 172.16.10.2/30 |
        | ic-A2 | conn-B          | LF6_A  | 172.16.10.5/30 |
        | ic-C2 | conn-B          | LF6_C  | 172.16.10.6/30 |

- **输出：**

    | ic_uuid | local_device | local_port | local_ip       | peer_device | peer_port  | peer_ip        |
    |---------|--------------|------------|----------------|-------------|------------|----------------|
    | ic-A1   | LF6_A        | 100GE1/0/2 | 172.16.10.1/30 | LF6_C       | 100GE1/0/2 | 172.16.10.2/30 |
    | ic-A2   | LF6_A        | 100GE1/0/3 | 172.16.10.5/30 | LF6_C       | 100GE1/0/3 | 172.16.10.6/30 |
    | ic-C1   | LF6_C        | 100GE1/0/2 | 172.16.10.2/30 | LF6_A       | 100GE1/0/2 | 172.16.10.1/30 |
    | ic-C2   | LF6_C        | 100GE1/0/3 | 172.16.10.6/30 | LF6_A       | 100GE1/0/3 | 172.16.10.5/30 |

- **解释：**

    - ic-A1 属于 conn-A，在 conn-A 下找设备不是 LF6_A 的那行，找到 ic-C1，取其 ip_address = 172.16.10.2/30 作为 peer_ip

    - ic-A2 属于 conn-B，同理找到 ic-C2，peer_ip = 172.16.10.6/30

    - ic-C1、ic-C2 各自反向配对，逻辑相同

### 示例 3：孤儿端——对端行缺失

- **输入：**

    - **`connections`：**

        | connection_uuid | from_device | from_port | to_device | to_port   |
        |-----------------|-------------|-----------|-----------|-----------|
        | conn-X          | DEV-A       | 10GE1/0/1 | DEV-B     | 10GE1/0/1 |

    - **`interconnects`：**（DEV-B 那行未录入）

        | uuid | connection_uuid | device | ip_address     |
        |------|-----------------|--------|----------------|
        | ic-A | conn-X          | DEV-A  | 192.168.0.1/30 |

- **输出：**

    | ic_uuid | local_device | local_port | local_ip       | peer_device | peer_port | peer_ip |
    |---------|--------------|------------|----------------|-------------|-----------|---------|
    | ic-A    | DEV-A        | 10GE1/0/1  | 192.168.0.1/30 | DEV-B       | 10GE1/0/1 | NULL    |

- **解释：**

    - **对端设备/接口：** 来自 `connections`，即使对端行缺失也能填

    - **对端 IP：** 来自配对行，interconnects 里找不到 DEV-B 那行，配不到则输出 NULL

    - **本端行：** 不能因为配不到对端而丢失，必须保留

