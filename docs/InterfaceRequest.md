# InterfaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | [**NestedDeviceRequest**](NestedDeviceRequest.md) |  | 
**Vdcs** | Pointer to **[]int32** |  | [optional] 
**Module** | Pointer to [**NullableComponentNestedModuleRequest**](ComponentNestedModuleRequest.md) |  | [optional] 
**Name** | **string** |  | 
**Label** | Pointer to **string** | Physical label | [optional] 
**Type** | **string** | * &#x60;virtual&#x60; - Virtual * &#x60;bridge&#x60; - Bridge * &#x60;lag&#x60; - Link Aggregation Group (LAG) * &#x60;100base-fx&#x60; - 100BASE-FX (10/100ME FIBER) * &#x60;100base-lfx&#x60; - 100BASE-LFX (10/100ME FIBER) * &#x60;100base-tx&#x60; - 100BASE-TX (10/100ME) * &#x60;100base-t1&#x60; - 100BASE-T1 (10/100ME Single Pair) * &#x60;1000base-t&#x60; - 1000BASE-T (1GE) * &#x60;2.5gbase-t&#x60; - 2.5GBASE-T (2.5GE) * &#x60;5gbase-t&#x60; - 5GBASE-T (5GE) * &#x60;10gbase-t&#x60; - 10GBASE-T (10GE) * &#x60;10gbase-cx4&#x60; - 10GBASE-CX4 (10GE) * &#x60;1000base-x-gbic&#x60; - GBIC (1GE) * &#x60;1000base-x-sfp&#x60; - SFP (1GE) * &#x60;10gbase-x-sfpp&#x60; - SFP+ (10GE) * &#x60;10gbase-x-xfp&#x60; - XFP (10GE) * &#x60;10gbase-x-xenpak&#x60; - XENPAK (10GE) * &#x60;10gbase-x-x2&#x60; - X2 (10GE) * &#x60;25gbase-x-sfp28&#x60; - SFP28 (25GE) * &#x60;50gbase-x-sfp56&#x60; - SFP56 (50GE) * &#x60;40gbase-x-qsfpp&#x60; - QSFP+ (40GE) * &#x60;50gbase-x-sfp28&#x60; - QSFP28 (50GE) * &#x60;100gbase-x-cfp&#x60; - CFP (100GE) * &#x60;100gbase-x-cfp2&#x60; - CFP2 (100GE) * &#x60;200gbase-x-cfp2&#x60; - CFP2 (200GE) * &#x60;400gbase-x-cfp2&#x60; - CFP2 (400GE) * &#x60;100gbase-x-cfp4&#x60; - CFP4 (100GE) * &#x60;100gbase-x-cxp&#x60; - CXP (100GE) * &#x60;100gbase-x-cpak&#x60; - Cisco CPAK (100GE) * &#x60;100gbase-x-dsfp&#x60; - DSFP (100GE) * &#x60;100gbase-x-sfpdd&#x60; - SFP-DD (100GE) * &#x60;100gbase-x-qsfp28&#x60; - QSFP28 (100GE) * &#x60;100gbase-x-qsfpdd&#x60; - QSFP-DD (100GE) * &#x60;200gbase-x-qsfp56&#x60; - QSFP56 (200GE) * &#x60;200gbase-x-qsfpdd&#x60; - QSFP-DD (200GE) * &#x60;400gbase-x-qsfpdd&#x60; - QSFP-DD (400GE) * &#x60;400gbase-x-osfp&#x60; - OSFP (400GE) * &#x60;400gbase-x-cdfp&#x60; - CDFP (400GE) * &#x60;400gbase-x-cfp8&#x60; - CPF8 (400GE) * &#x60;800gbase-x-qsfpdd&#x60; - QSFP-DD (800GE) * &#x60;800gbase-x-osfp&#x60; - OSFP (800GE) * &#x60;1000base-kx&#x60; - 1000BASE-KX (1GE) * &#x60;10gbase-kr&#x60; - 10GBASE-KR (10GE) * &#x60;10gbase-kx4&#x60; - 10GBASE-KX4 (10GE) * &#x60;25gbase-kr&#x60; - 25GBASE-KR (25GE) * &#x60;40gbase-kr4&#x60; - 40GBASE-KR4 (40GE) * &#x60;50gbase-kr&#x60; - 50GBASE-KR (50GE) * &#x60;100gbase-kp4&#x60; - 100GBASE-KP4 (100GE) * &#x60;100gbase-kr2&#x60; - 100GBASE-KR2 (100GE) * &#x60;100gbase-kr4&#x60; - 100GBASE-KR4 (100GE) * &#x60;ieee802.11a&#x60; - IEEE 802.11a * &#x60;ieee802.11g&#x60; - IEEE 802.11b/g * &#x60;ieee802.11n&#x60; - IEEE 802.11n * &#x60;ieee802.11ac&#x60; - IEEE 802.11ac * &#x60;ieee802.11ad&#x60; - IEEE 802.11ad * &#x60;ieee802.11ax&#x60; - IEEE 802.11ax * &#x60;ieee802.11ay&#x60; - IEEE 802.11ay * &#x60;ieee802.15.1&#x60; - IEEE 802.15.1 (Bluetooth) * &#x60;other-wireless&#x60; - Other (Wireless) * &#x60;gsm&#x60; - GSM * &#x60;cdma&#x60; - CDMA * &#x60;lte&#x60; - LTE * &#x60;sonet-oc3&#x60; - OC-3/STM-1 * &#x60;sonet-oc12&#x60; - OC-12/STM-4 * &#x60;sonet-oc48&#x60; - OC-48/STM-16 * &#x60;sonet-oc192&#x60; - OC-192/STM-64 * &#x60;sonet-oc768&#x60; - OC-768/STM-256 * &#x60;sonet-oc1920&#x60; - OC-1920/STM-640 * &#x60;sonet-oc3840&#x60; - OC-3840/STM-1234 * &#x60;1gfc-sfp&#x60; - SFP (1GFC) * &#x60;2gfc-sfp&#x60; - SFP (2GFC) * &#x60;4gfc-sfp&#x60; - SFP (4GFC) * &#x60;8gfc-sfpp&#x60; - SFP+ (8GFC) * &#x60;16gfc-sfpp&#x60; - SFP+ (16GFC) * &#x60;32gfc-sfp28&#x60; - SFP28 (32GFC) * &#x60;64gfc-qsfpp&#x60; - QSFP+ (64GFC) * &#x60;128gfc-qsfp28&#x60; - QSFP28 (128GFC) * &#x60;infiniband-sdr&#x60; - SDR (2 Gbps) * &#x60;infiniband-ddr&#x60; - DDR (4 Gbps) * &#x60;infiniband-qdr&#x60; - QDR (8 Gbps) * &#x60;infiniband-fdr10&#x60; - FDR10 (10 Gbps) * &#x60;infiniband-fdr&#x60; - FDR (13.5 Gbps) * &#x60;infiniband-edr&#x60; - EDR (25 Gbps) * &#x60;infiniband-hdr&#x60; - HDR (50 Gbps) * &#x60;infiniband-ndr&#x60; - NDR (100 Gbps) * &#x60;infiniband-xdr&#x60; - XDR (250 Gbps) * &#x60;t1&#x60; - T1 (1.544 Mbps) * &#x60;e1&#x60; - E1 (2.048 Mbps) * &#x60;t3&#x60; - T3 (45 Mbps) * &#x60;e3&#x60; - E3 (34 Mbps) * &#x60;xdsl&#x60; - xDSL * &#x60;docsis&#x60; - DOCSIS * &#x60;gpon&#x60; - GPON (2.5 Gbps / 1.25 Gps) * &#x60;xg-pon&#x60; - XG-PON (10 Gbps / 2.5 Gbps) * &#x60;xgs-pon&#x60; - XGS-PON (10 Gbps) * &#x60;ng-pon2&#x60; - NG-PON2 (TWDM-PON) (4x10 Gbps) * &#x60;epon&#x60; - EPON (1 Gbps) * &#x60;10g-epon&#x60; - 10G-EPON (10 Gbps) * &#x60;cisco-stackwise&#x60; - Cisco StackWise * &#x60;cisco-stackwise-plus&#x60; - Cisco StackWise Plus * &#x60;cisco-flexstack&#x60; - Cisco FlexStack * &#x60;cisco-flexstack-plus&#x60; - Cisco FlexStack Plus * &#x60;cisco-stackwise-80&#x60; - Cisco StackWise-80 * &#x60;cisco-stackwise-160&#x60; - Cisco StackWise-160 * &#x60;cisco-stackwise-320&#x60; - Cisco StackWise-320 * &#x60;cisco-stackwise-480&#x60; - Cisco StackWise-480 * &#x60;cisco-stackwise-1t&#x60; - Cisco StackWise-1T * &#x60;juniper-vcp&#x60; - Juniper VCP * &#x60;extreme-summitstack&#x60; - Extreme SummitStack * &#x60;extreme-summitstack-128&#x60; - Extreme SummitStack-128 * &#x60;extreme-summitstack-256&#x60; - Extreme SummitStack-256 * &#x60;extreme-summitstack-512&#x60; - Extreme SummitStack-512 * &#x60;other&#x60; - Other | 
**Enabled** | Pointer to **bool** |  | [optional] 
**Parent** | Pointer to [**NullableNestedInterfaceRequest**](NestedInterfaceRequest.md) |  | [optional] 
**Bridge** | Pointer to [**NullableNestedInterfaceRequest**](NestedInterfaceRequest.md) |  | [optional] 
**Lag** | Pointer to [**NullableNestedInterfaceRequest**](NestedInterfaceRequest.md) |  | [optional] 
**Mtu** | Pointer to **NullableInt32** |  | [optional] 
**MacAddress** | Pointer to **NullableString** |  | [optional] 
**Speed** | Pointer to **NullableInt32** |  | [optional] 
**Duplex** | Pointer to **NullableString** | * &#x60;half&#x60; - Half * &#x60;full&#x60; - Full * &#x60;auto&#x60; - Auto | [optional] 
**Wwn** | Pointer to **NullableString** |  | [optional] 
**MgmtOnly** | Pointer to **bool** | This interface is used only for out-of-band management | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to **string** | * &#x60;access&#x60; - Access * &#x60;tagged&#x60; - Tagged * &#x60;tagged-all&#x60; - Tagged (All) | [optional] 
**RfRole** | Pointer to **string** | * &#x60;ap&#x60; - Access point * &#x60;station&#x60; - Station | [optional] 
**RfChannel** | Pointer to **string** | * &#x60;2.4g-1-2412-22&#x60; - 1 (2412 MHz) * &#x60;2.4g-2-2417-22&#x60; - 2 (2417 MHz) * &#x60;2.4g-3-2422-22&#x60; - 3 (2422 MHz) * &#x60;2.4g-4-2427-22&#x60; - 4 (2427 MHz) * &#x60;2.4g-5-2432-22&#x60; - 5 (2432 MHz) * &#x60;2.4g-6-2437-22&#x60; - 6 (2437 MHz) * &#x60;2.4g-7-2442-22&#x60; - 7 (2442 MHz) * &#x60;2.4g-8-2447-22&#x60; - 8 (2447 MHz) * &#x60;2.4g-9-2452-22&#x60; - 9 (2452 MHz) * &#x60;2.4g-10-2457-22&#x60; - 10 (2457 MHz) * &#x60;2.4g-11-2462-22&#x60; - 11 (2462 MHz) * &#x60;2.4g-12-2467-22&#x60; - 12 (2467 MHz) * &#x60;2.4g-13-2472-22&#x60; - 13 (2472 MHz) * &#x60;5g-32-5160-20&#x60; - 32 (5160/20 MHz) * &#x60;5g-34-5170-40&#x60; - 34 (5170/40 MHz) * &#x60;5g-36-5180-20&#x60; - 36 (5180/20 MHz) * &#x60;5g-38-5190-40&#x60; - 38 (5190/40 MHz) * &#x60;5g-40-5200-20&#x60; - 40 (5200/20 MHz) * &#x60;5g-42-5210-80&#x60; - 42 (5210/80 MHz) * &#x60;5g-44-5220-20&#x60; - 44 (5220/20 MHz) * &#x60;5g-46-5230-40&#x60; - 46 (5230/40 MHz) * &#x60;5g-48-5240-20&#x60; - 48 (5240/20 MHz) * &#x60;5g-50-5250-160&#x60; - 50 (5250/160 MHz) * &#x60;5g-52-5260-20&#x60; - 52 (5260/20 MHz) * &#x60;5g-54-5270-40&#x60; - 54 (5270/40 MHz) * &#x60;5g-56-5280-20&#x60; - 56 (5280/20 MHz) * &#x60;5g-58-5290-80&#x60; - 58 (5290/80 MHz) * &#x60;5g-60-5300-20&#x60; - 60 (5300/20 MHz) * &#x60;5g-62-5310-40&#x60; - 62 (5310/40 MHz) * &#x60;5g-64-5320-20&#x60; - 64 (5320/20 MHz) * &#x60;5g-100-5500-20&#x60; - 100 (5500/20 MHz) * &#x60;5g-102-5510-40&#x60; - 102 (5510/40 MHz) * &#x60;5g-104-5520-20&#x60; - 104 (5520/20 MHz) * &#x60;5g-106-5530-80&#x60; - 106 (5530/80 MHz) * &#x60;5g-108-5540-20&#x60; - 108 (5540/20 MHz) * &#x60;5g-110-5550-40&#x60; - 110 (5550/40 MHz) * &#x60;5g-112-5560-20&#x60; - 112 (5560/20 MHz) * &#x60;5g-114-5570-160&#x60; - 114 (5570/160 MHz) * &#x60;5g-116-5580-20&#x60; - 116 (5580/20 MHz) * &#x60;5g-118-5590-40&#x60; - 118 (5590/40 MHz) * &#x60;5g-120-5600-20&#x60; - 120 (5600/20 MHz) * &#x60;5g-122-5610-80&#x60; - 122 (5610/80 MHz) * &#x60;5g-124-5620-20&#x60; - 124 (5620/20 MHz) * &#x60;5g-126-5630-40&#x60; - 126 (5630/40 MHz) * &#x60;5g-128-5640-20&#x60; - 128 (5640/20 MHz) * &#x60;5g-132-5660-20&#x60; - 132 (5660/20 MHz) * &#x60;5g-134-5670-40&#x60; - 134 (5670/40 MHz) * &#x60;5g-136-5680-20&#x60; - 136 (5680/20 MHz) * &#x60;5g-138-5690-80&#x60; - 138 (5690/80 MHz) * &#x60;5g-140-5700-20&#x60; - 140 (5700/20 MHz) * &#x60;5g-142-5710-40&#x60; - 142 (5710/40 MHz) * &#x60;5g-144-5720-20&#x60; - 144 (5720/20 MHz) * &#x60;5g-149-5745-20&#x60; - 149 (5745/20 MHz) * &#x60;5g-151-5755-40&#x60; - 151 (5755/40 MHz) * &#x60;5g-153-5765-20&#x60; - 153 (5765/20 MHz) * &#x60;5g-155-5775-80&#x60; - 155 (5775/80 MHz) * &#x60;5g-157-5785-20&#x60; - 157 (5785/20 MHz) * &#x60;5g-159-5795-40&#x60; - 159 (5795/40 MHz) * &#x60;5g-161-5805-20&#x60; - 161 (5805/20 MHz) * &#x60;5g-163-5815-160&#x60; - 163 (5815/160 MHz) * &#x60;5g-165-5825-20&#x60; - 165 (5825/20 MHz) * &#x60;5g-167-5835-40&#x60; - 167 (5835/40 MHz) * &#x60;5g-169-5845-20&#x60; - 169 (5845/20 MHz) * &#x60;5g-171-5855-80&#x60; - 171 (5855/80 MHz) * &#x60;5g-173-5865-20&#x60; - 173 (5865/20 MHz) * &#x60;5g-175-5875-40&#x60; - 175 (5875/40 MHz) * &#x60;5g-177-5885-20&#x60; - 177 (5885/20 MHz) * &#x60;6g-1-5955-20&#x60; - 1 (5955/20 MHz) * &#x60;6g-3-5965-40&#x60; - 3 (5965/40 MHz) * &#x60;6g-5-5975-20&#x60; - 5 (5975/20 MHz) * &#x60;6g-7-5985-80&#x60; - 7 (5985/80 MHz) * &#x60;6g-9-5995-20&#x60; - 9 (5995/20 MHz) * &#x60;6g-11-6005-40&#x60; - 11 (6005/40 MHz) * &#x60;6g-13-6015-20&#x60; - 13 (6015/20 MHz) * &#x60;6g-15-6025-160&#x60; - 15 (6025/160 MHz) * &#x60;6g-17-6035-20&#x60; - 17 (6035/20 MHz) * &#x60;6g-19-6045-40&#x60; - 19 (6045/40 MHz) * &#x60;6g-21-6055-20&#x60; - 21 (6055/20 MHz) * &#x60;6g-23-6065-80&#x60; - 23 (6065/80 MHz) * &#x60;6g-25-6075-20&#x60; - 25 (6075/20 MHz) * &#x60;6g-27-6085-40&#x60; - 27 (6085/40 MHz) * &#x60;6g-29-6095-20&#x60; - 29 (6095/20 MHz) * &#x60;6g-31-6105-320&#x60; - 31 (6105/320 MHz) * &#x60;6g-33-6115-20&#x60; - 33 (6115/20 MHz) * &#x60;6g-35-6125-40&#x60; - 35 (6125/40 MHz) * &#x60;6g-37-6135-20&#x60; - 37 (6135/20 MHz) * &#x60;6g-39-6145-80&#x60; - 39 (6145/80 MHz) * &#x60;6g-41-6155-20&#x60; - 41 (6155/20 MHz) * &#x60;6g-43-6165-40&#x60; - 43 (6165/40 MHz) * &#x60;6g-45-6175-20&#x60; - 45 (6175/20 MHz) * &#x60;6g-47-6185-160&#x60; - 47 (6185/160 MHz) * &#x60;6g-49-6195-20&#x60; - 49 (6195/20 MHz) * &#x60;6g-51-6205-40&#x60; - 51 (6205/40 MHz) * &#x60;6g-53-6215-20&#x60; - 53 (6215/20 MHz) * &#x60;6g-55-6225-80&#x60; - 55 (6225/80 MHz) * &#x60;6g-57-6235-20&#x60; - 57 (6235/20 MHz) * &#x60;6g-59-6245-40&#x60; - 59 (6245/40 MHz) * &#x60;6g-61-6255-20&#x60; - 61 (6255/20 MHz) * &#x60;6g-65-6275-20&#x60; - 65 (6275/20 MHz) * &#x60;6g-67-6285-40&#x60; - 67 (6285/40 MHz) * &#x60;6g-69-6295-20&#x60; - 69 (6295/20 MHz) * &#x60;6g-71-6305-80&#x60; - 71 (6305/80 MHz) * &#x60;6g-73-6315-20&#x60; - 73 (6315/20 MHz) * &#x60;6g-75-6325-40&#x60; - 75 (6325/40 MHz) * &#x60;6g-77-6335-20&#x60; - 77 (6335/20 MHz) * &#x60;6g-79-6345-160&#x60; - 79 (6345/160 MHz) * &#x60;6g-81-6355-20&#x60; - 81 (6355/20 MHz) * &#x60;6g-83-6365-40&#x60; - 83 (6365/40 MHz) * &#x60;6g-85-6375-20&#x60; - 85 (6375/20 MHz) * &#x60;6g-87-6385-80&#x60; - 87 (6385/80 MHz) * &#x60;6g-89-6395-20&#x60; - 89 (6395/20 MHz) * &#x60;6g-91-6405-40&#x60; - 91 (6405/40 MHz) * &#x60;6g-93-6415-20&#x60; - 93 (6415/20 MHz) * &#x60;6g-95-6425-320&#x60; - 95 (6425/320 MHz) * &#x60;6g-97-6435-20&#x60; - 97 (6435/20 MHz) * &#x60;6g-99-6445-40&#x60; - 99 (6445/40 MHz) * &#x60;6g-101-6455-20&#x60; - 101 (6455/20 MHz) * &#x60;6g-103-6465-80&#x60; - 103 (6465/80 MHz) * &#x60;6g-105-6475-20&#x60; - 105 (6475/20 MHz) * &#x60;6g-107-6485-40&#x60; - 107 (6485/40 MHz) * &#x60;6g-109-6495-20&#x60; - 109 (6495/20 MHz) * &#x60;6g-111-6505-160&#x60; - 111 (6505/160 MHz) * &#x60;6g-113-6515-20&#x60; - 113 (6515/20 MHz) * &#x60;6g-115-6525-40&#x60; - 115 (6525/40 MHz) * &#x60;6g-117-6535-20&#x60; - 117 (6535/20 MHz) * &#x60;6g-119-6545-80&#x60; - 119 (6545/80 MHz) * &#x60;6g-121-6555-20&#x60; - 121 (6555/20 MHz) * &#x60;6g-123-6565-40&#x60; - 123 (6565/40 MHz) * &#x60;6g-125-6575-20&#x60; - 125 (6575/20 MHz) * &#x60;6g-129-6595-20&#x60; - 129 (6595/20 MHz) * &#x60;6g-131-6605-40&#x60; - 131 (6605/40 MHz) * &#x60;6g-133-6615-20&#x60; - 133 (6615/20 MHz) * &#x60;6g-135-6625-80&#x60; - 135 (6625/80 MHz) * &#x60;6g-137-6635-20&#x60; - 137 (6635/20 MHz) * &#x60;6g-139-6645-40&#x60; - 139 (6645/40 MHz) * &#x60;6g-141-6655-20&#x60; - 141 (6655/20 MHz) * &#x60;6g-143-6665-160&#x60; - 143 (6665/160 MHz) * &#x60;6g-145-6675-20&#x60; - 145 (6675/20 MHz) * &#x60;6g-147-6685-40&#x60; - 147 (6685/40 MHz) * &#x60;6g-149-6695-20&#x60; - 149 (6695/20 MHz) * &#x60;6g-151-6705-80&#x60; - 151 (6705/80 MHz) * &#x60;6g-153-6715-20&#x60; - 153 (6715/20 MHz) * &#x60;6g-155-6725-40&#x60; - 155 (6725/40 MHz) * &#x60;6g-157-6735-20&#x60; - 157 (6735/20 MHz) * &#x60;6g-159-6745-320&#x60; - 159 (6745/320 MHz) * &#x60;6g-161-6755-20&#x60; - 161 (6755/20 MHz) * &#x60;6g-163-6765-40&#x60; - 163 (6765/40 MHz) * &#x60;6g-165-6775-20&#x60; - 165 (6775/20 MHz) * &#x60;6g-167-6785-80&#x60; - 167 (6785/80 MHz) * &#x60;6g-169-6795-20&#x60; - 169 (6795/20 MHz) * &#x60;6g-171-6805-40&#x60; - 171 (6805/40 MHz) * &#x60;6g-173-6815-20&#x60; - 173 (6815/20 MHz) * &#x60;6g-175-6825-160&#x60; - 175 (6825/160 MHz) * &#x60;6g-177-6835-20&#x60; - 177 (6835/20 MHz) * &#x60;6g-179-6845-40&#x60; - 179 (6845/40 MHz) * &#x60;6g-181-6855-20&#x60; - 181 (6855/20 MHz) * &#x60;6g-183-6865-80&#x60; - 183 (6865/80 MHz) * &#x60;6g-185-6875-20&#x60; - 185 (6875/20 MHz) * &#x60;6g-187-6885-40&#x60; - 187 (6885/40 MHz) * &#x60;6g-189-6895-20&#x60; - 189 (6895/20 MHz) * &#x60;6g-193-6915-20&#x60; - 193 (6915/20 MHz) * &#x60;6g-195-6925-40&#x60; - 195 (6925/40 MHz) * &#x60;6g-197-6935-20&#x60; - 197 (6935/20 MHz) * &#x60;6g-199-6945-80&#x60; - 199 (6945/80 MHz) * &#x60;6g-201-6955-20&#x60; - 201 (6955/20 MHz) * &#x60;6g-203-6965-40&#x60; - 203 (6965/40 MHz) * &#x60;6g-205-6975-20&#x60; - 205 (6975/20 MHz) * &#x60;6g-207-6985-160&#x60; - 207 (6985/160 MHz) * &#x60;6g-209-6995-20&#x60; - 209 (6995/20 MHz) * &#x60;6g-211-7005-40&#x60; - 211 (7005/40 MHz) * &#x60;6g-213-7015-20&#x60; - 213 (7015/20 MHz) * &#x60;6g-215-7025-80&#x60; - 215 (7025/80 MHz) * &#x60;6g-217-7035-20&#x60; - 217 (7035/20 MHz) * &#x60;6g-219-7045-40&#x60; - 219 (7045/40 MHz) * &#x60;6g-221-7055-20&#x60; - 221 (7055/20 MHz) * &#x60;6g-225-7075-20&#x60; - 225 (7075/20 MHz) * &#x60;6g-227-7085-40&#x60; - 227 (7085/40 MHz) * &#x60;6g-229-7095-20&#x60; - 229 (7095/20 MHz) * &#x60;6g-233-7115-20&#x60; - 233 (7115/20 MHz) * &#x60;60g-1-58320-2160&#x60; - 1 (58.32/2.16 GHz) * &#x60;60g-2-60480-2160&#x60; - 2 (60.48/2.16 GHz) * &#x60;60g-3-62640-2160&#x60; - 3 (62.64/2.16 GHz) * &#x60;60g-4-64800-2160&#x60; - 4 (64.80/2.16 GHz) * &#x60;60g-5-66960-2160&#x60; - 5 (66.96/2.16 GHz) * &#x60;60g-6-69120-2160&#x60; - 6 (69.12/2.16 GHz) * &#x60;60g-9-59400-4320&#x60; - 9 (59.40/4.32 GHz) * &#x60;60g-10-61560-4320&#x60; - 10 (61.56/4.32 GHz) * &#x60;60g-11-63720-4320&#x60; - 11 (63.72/4.32 GHz) * &#x60;60g-12-65880-4320&#x60; - 12 (65.88/4.32 GHz) * &#x60;60g-13-68040-4320&#x60; - 13 (68.04/4.32 GHz) * &#x60;60g-17-60480-6480&#x60; - 17 (60.48/6.48 GHz) * &#x60;60g-18-62640-6480&#x60; - 18 (62.64/6.48 GHz) * &#x60;60g-19-64800-6480&#x60; - 19 (64.80/6.48 GHz) * &#x60;60g-20-66960-6480&#x60; - 20 (66.96/6.48 GHz) * &#x60;60g-25-61560-6480&#x60; - 25 (61.56/8.64 GHz) * &#x60;60g-26-63720-6480&#x60; - 26 (63.72/8.64 GHz) * &#x60;60g-27-65880-6480&#x60; - 27 (65.88/8.64 GHz) | [optional] 
**PoeMode** | Pointer to **string** | * &#x60;pd&#x60; - PD * &#x60;pse&#x60; - PSE | [optional] 
**PoeType** | Pointer to **string** | * &#x60;type1-ieee802.3af&#x60; - 802.3af (Type 1) * &#x60;type2-ieee802.3at&#x60; - 802.3at (Type 2) * &#x60;type3-ieee802.3bt&#x60; - 802.3bt (Type 3) * &#x60;type4-ieee802.3bt&#x60; - 802.3bt (Type 4) * &#x60;passive-24v-2pair&#x60; - Passive 24V (2-pair) * &#x60;passive-24v-4pair&#x60; - Passive 24V (4-pair) * &#x60;passive-48v-2pair&#x60; - Passive 48V (2-pair) * &#x60;passive-48v-4pair&#x60; - Passive 48V (4-pair) | [optional] 
**RfChannelFrequency** | Pointer to **NullableFloat64** | Populated by selected channel (if set) | [optional] 
**RfChannelWidth** | Pointer to **NullableFloat64** | Populated by selected channel (if set) | [optional] 
**TxPower** | Pointer to **NullableInt32** |  | [optional] 
**UntaggedVlan** | Pointer to [**NullableNestedVLANRequest**](NestedVLANRequest.md) |  | [optional] 
**TaggedVlans** | Pointer to **[]int32** |  | [optional] 
**MarkConnected** | Pointer to **bool** | Treat as if a cable is connected | [optional] 
**WirelessLans** | Pointer to **[]int32** |  | [optional] 
**Vrf** | Pointer to [**NullableNestedVRFRequest**](NestedVRFRequest.md) |  | [optional] 
**Tags** | Pointer to [**[]NestedTagRequest**](NestedTagRequest.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewInterfaceRequest

`func NewInterfaceRequest(device NestedDeviceRequest, name string, type_ string, ) *InterfaceRequest`

NewInterfaceRequest instantiates a new InterfaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterfaceRequestWithDefaults

`func NewInterfaceRequestWithDefaults() *InterfaceRequest`

NewInterfaceRequestWithDefaults instantiates a new InterfaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *InterfaceRequest) GetDevice() NestedDeviceRequest`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *InterfaceRequest) GetDeviceOk() (*NestedDeviceRequest, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *InterfaceRequest) SetDevice(v NestedDeviceRequest)`

SetDevice sets Device field to given value.


### GetVdcs

`func (o *InterfaceRequest) GetVdcs() []int32`

GetVdcs returns the Vdcs field if non-nil, zero value otherwise.

### GetVdcsOk

`func (o *InterfaceRequest) GetVdcsOk() (*[]int32, bool)`

GetVdcsOk returns a tuple with the Vdcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdcs

`func (o *InterfaceRequest) SetVdcs(v []int32)`

SetVdcs sets Vdcs field to given value.

### HasVdcs

`func (o *InterfaceRequest) HasVdcs() bool`

HasVdcs returns a boolean if a field has been set.

### GetModule

`func (o *InterfaceRequest) GetModule() ComponentNestedModuleRequest`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *InterfaceRequest) GetModuleOk() (*ComponentNestedModuleRequest, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *InterfaceRequest) SetModule(v ComponentNestedModuleRequest)`

SetModule sets Module field to given value.

### HasModule

`func (o *InterfaceRequest) HasModule() bool`

HasModule returns a boolean if a field has been set.

### SetModuleNil

`func (o *InterfaceRequest) SetModuleNil(b bool)`

 SetModuleNil sets the value for Module to be an explicit nil

### UnsetModule
`func (o *InterfaceRequest) UnsetModule()`

UnsetModule ensures that no value is present for Module, not even an explicit nil
### GetName

`func (o *InterfaceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InterfaceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InterfaceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *InterfaceRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *InterfaceRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *InterfaceRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *InterfaceRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *InterfaceRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InterfaceRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InterfaceRequest) SetType(v string)`

SetType sets Type field to given value.


### GetEnabled

`func (o *InterfaceRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InterfaceRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InterfaceRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InterfaceRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetParent

`func (o *InterfaceRequest) GetParent() NestedInterfaceRequest`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *InterfaceRequest) GetParentOk() (*NestedInterfaceRequest, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *InterfaceRequest) SetParent(v NestedInterfaceRequest)`

SetParent sets Parent field to given value.

### HasParent

`func (o *InterfaceRequest) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *InterfaceRequest) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *InterfaceRequest) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetBridge

`func (o *InterfaceRequest) GetBridge() NestedInterfaceRequest`

GetBridge returns the Bridge field if non-nil, zero value otherwise.

### GetBridgeOk

`func (o *InterfaceRequest) GetBridgeOk() (*NestedInterfaceRequest, bool)`

GetBridgeOk returns a tuple with the Bridge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridge

`func (o *InterfaceRequest) SetBridge(v NestedInterfaceRequest)`

SetBridge sets Bridge field to given value.

### HasBridge

`func (o *InterfaceRequest) HasBridge() bool`

HasBridge returns a boolean if a field has been set.

### SetBridgeNil

`func (o *InterfaceRequest) SetBridgeNil(b bool)`

 SetBridgeNil sets the value for Bridge to be an explicit nil

### UnsetBridge
`func (o *InterfaceRequest) UnsetBridge()`

UnsetBridge ensures that no value is present for Bridge, not even an explicit nil
### GetLag

`func (o *InterfaceRequest) GetLag() NestedInterfaceRequest`

GetLag returns the Lag field if non-nil, zero value otherwise.

### GetLagOk

`func (o *InterfaceRequest) GetLagOk() (*NestedInterfaceRequest, bool)`

GetLagOk returns a tuple with the Lag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLag

`func (o *InterfaceRequest) SetLag(v NestedInterfaceRequest)`

SetLag sets Lag field to given value.

### HasLag

`func (o *InterfaceRequest) HasLag() bool`

HasLag returns a boolean if a field has been set.

### SetLagNil

`func (o *InterfaceRequest) SetLagNil(b bool)`

 SetLagNil sets the value for Lag to be an explicit nil

### UnsetLag
`func (o *InterfaceRequest) UnsetLag()`

UnsetLag ensures that no value is present for Lag, not even an explicit nil
### GetMtu

`func (o *InterfaceRequest) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *InterfaceRequest) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *InterfaceRequest) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *InterfaceRequest) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### SetMtuNil

`func (o *InterfaceRequest) SetMtuNil(b bool)`

 SetMtuNil sets the value for Mtu to be an explicit nil

### UnsetMtu
`func (o *InterfaceRequest) UnsetMtu()`

UnsetMtu ensures that no value is present for Mtu, not even an explicit nil
### GetMacAddress

`func (o *InterfaceRequest) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *InterfaceRequest) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *InterfaceRequest) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *InterfaceRequest) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### SetMacAddressNil

`func (o *InterfaceRequest) SetMacAddressNil(b bool)`

 SetMacAddressNil sets the value for MacAddress to be an explicit nil

### UnsetMacAddress
`func (o *InterfaceRequest) UnsetMacAddress()`

UnsetMacAddress ensures that no value is present for MacAddress, not even an explicit nil
### GetSpeed

`func (o *InterfaceRequest) GetSpeed() int32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *InterfaceRequest) GetSpeedOk() (*int32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *InterfaceRequest) SetSpeed(v int32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *InterfaceRequest) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### SetSpeedNil

`func (o *InterfaceRequest) SetSpeedNil(b bool)`

 SetSpeedNil sets the value for Speed to be an explicit nil

### UnsetSpeed
`func (o *InterfaceRequest) UnsetSpeed()`

UnsetSpeed ensures that no value is present for Speed, not even an explicit nil
### GetDuplex

`func (o *InterfaceRequest) GetDuplex() string`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *InterfaceRequest) GetDuplexOk() (*string, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *InterfaceRequest) SetDuplex(v string)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *InterfaceRequest) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### SetDuplexNil

`func (o *InterfaceRequest) SetDuplexNil(b bool)`

 SetDuplexNil sets the value for Duplex to be an explicit nil

### UnsetDuplex
`func (o *InterfaceRequest) UnsetDuplex()`

UnsetDuplex ensures that no value is present for Duplex, not even an explicit nil
### GetWwn

`func (o *InterfaceRequest) GetWwn() string`

GetWwn returns the Wwn field if non-nil, zero value otherwise.

### GetWwnOk

`func (o *InterfaceRequest) GetWwnOk() (*string, bool)`

GetWwnOk returns a tuple with the Wwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwn

`func (o *InterfaceRequest) SetWwn(v string)`

SetWwn sets Wwn field to given value.

### HasWwn

`func (o *InterfaceRequest) HasWwn() bool`

HasWwn returns a boolean if a field has been set.

### SetWwnNil

`func (o *InterfaceRequest) SetWwnNil(b bool)`

 SetWwnNil sets the value for Wwn to be an explicit nil

### UnsetWwn
`func (o *InterfaceRequest) UnsetWwn()`

UnsetWwn ensures that no value is present for Wwn, not even an explicit nil
### GetMgmtOnly

`func (o *InterfaceRequest) GetMgmtOnly() bool`

GetMgmtOnly returns the MgmtOnly field if non-nil, zero value otherwise.

### GetMgmtOnlyOk

`func (o *InterfaceRequest) GetMgmtOnlyOk() (*bool, bool)`

GetMgmtOnlyOk returns a tuple with the MgmtOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtOnly

`func (o *InterfaceRequest) SetMgmtOnly(v bool)`

SetMgmtOnly sets MgmtOnly field to given value.

### HasMgmtOnly

`func (o *InterfaceRequest) HasMgmtOnly() bool`

HasMgmtOnly returns a boolean if a field has been set.

### GetDescription

`func (o *InterfaceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InterfaceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InterfaceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InterfaceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMode

`func (o *InterfaceRequest) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InterfaceRequest) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InterfaceRequest) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InterfaceRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetRfRole

`func (o *InterfaceRequest) GetRfRole() string`

GetRfRole returns the RfRole field if non-nil, zero value otherwise.

### GetRfRoleOk

`func (o *InterfaceRequest) GetRfRoleOk() (*string, bool)`

GetRfRoleOk returns a tuple with the RfRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfRole

`func (o *InterfaceRequest) SetRfRole(v string)`

SetRfRole sets RfRole field to given value.

### HasRfRole

`func (o *InterfaceRequest) HasRfRole() bool`

HasRfRole returns a boolean if a field has been set.

### GetRfChannel

`func (o *InterfaceRequest) GetRfChannel() string`

GetRfChannel returns the RfChannel field if non-nil, zero value otherwise.

### GetRfChannelOk

`func (o *InterfaceRequest) GetRfChannelOk() (*string, bool)`

GetRfChannelOk returns a tuple with the RfChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfChannel

`func (o *InterfaceRequest) SetRfChannel(v string)`

SetRfChannel sets RfChannel field to given value.

### HasRfChannel

`func (o *InterfaceRequest) HasRfChannel() bool`

HasRfChannel returns a boolean if a field has been set.

### GetPoeMode

`func (o *InterfaceRequest) GetPoeMode() string`

GetPoeMode returns the PoeMode field if non-nil, zero value otherwise.

### GetPoeModeOk

`func (o *InterfaceRequest) GetPoeModeOk() (*string, bool)`

GetPoeModeOk returns a tuple with the PoeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeMode

`func (o *InterfaceRequest) SetPoeMode(v string)`

SetPoeMode sets PoeMode field to given value.

### HasPoeMode

`func (o *InterfaceRequest) HasPoeMode() bool`

HasPoeMode returns a boolean if a field has been set.

### GetPoeType

`func (o *InterfaceRequest) GetPoeType() string`

GetPoeType returns the PoeType field if non-nil, zero value otherwise.

### GetPoeTypeOk

`func (o *InterfaceRequest) GetPoeTypeOk() (*string, bool)`

GetPoeTypeOk returns a tuple with the PoeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoeType

`func (o *InterfaceRequest) SetPoeType(v string)`

SetPoeType sets PoeType field to given value.

### HasPoeType

`func (o *InterfaceRequest) HasPoeType() bool`

HasPoeType returns a boolean if a field has been set.

### GetRfChannelFrequency

`func (o *InterfaceRequest) GetRfChannelFrequency() float64`

GetRfChannelFrequency returns the RfChannelFrequency field if non-nil, zero value otherwise.

### GetRfChannelFrequencyOk

`func (o *InterfaceRequest) GetRfChannelFrequencyOk() (*float64, bool)`

GetRfChannelFrequencyOk returns a tuple with the RfChannelFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfChannelFrequency

`func (o *InterfaceRequest) SetRfChannelFrequency(v float64)`

SetRfChannelFrequency sets RfChannelFrequency field to given value.

### HasRfChannelFrequency

`func (o *InterfaceRequest) HasRfChannelFrequency() bool`

HasRfChannelFrequency returns a boolean if a field has been set.

### SetRfChannelFrequencyNil

`func (o *InterfaceRequest) SetRfChannelFrequencyNil(b bool)`

 SetRfChannelFrequencyNil sets the value for RfChannelFrequency to be an explicit nil

### UnsetRfChannelFrequency
`func (o *InterfaceRequest) UnsetRfChannelFrequency()`

UnsetRfChannelFrequency ensures that no value is present for RfChannelFrequency, not even an explicit nil
### GetRfChannelWidth

`func (o *InterfaceRequest) GetRfChannelWidth() float64`

GetRfChannelWidth returns the RfChannelWidth field if non-nil, zero value otherwise.

### GetRfChannelWidthOk

`func (o *InterfaceRequest) GetRfChannelWidthOk() (*float64, bool)`

GetRfChannelWidthOk returns a tuple with the RfChannelWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfChannelWidth

`func (o *InterfaceRequest) SetRfChannelWidth(v float64)`

SetRfChannelWidth sets RfChannelWidth field to given value.

### HasRfChannelWidth

`func (o *InterfaceRequest) HasRfChannelWidth() bool`

HasRfChannelWidth returns a boolean if a field has been set.

### SetRfChannelWidthNil

`func (o *InterfaceRequest) SetRfChannelWidthNil(b bool)`

 SetRfChannelWidthNil sets the value for RfChannelWidth to be an explicit nil

### UnsetRfChannelWidth
`func (o *InterfaceRequest) UnsetRfChannelWidth()`

UnsetRfChannelWidth ensures that no value is present for RfChannelWidth, not even an explicit nil
### GetTxPower

`func (o *InterfaceRequest) GetTxPower() int32`

GetTxPower returns the TxPower field if non-nil, zero value otherwise.

### GetTxPowerOk

`func (o *InterfaceRequest) GetTxPowerOk() (*int32, bool)`

GetTxPowerOk returns a tuple with the TxPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxPower

`func (o *InterfaceRequest) SetTxPower(v int32)`

SetTxPower sets TxPower field to given value.

### HasTxPower

`func (o *InterfaceRequest) HasTxPower() bool`

HasTxPower returns a boolean if a field has been set.

### SetTxPowerNil

`func (o *InterfaceRequest) SetTxPowerNil(b bool)`

 SetTxPowerNil sets the value for TxPower to be an explicit nil

### UnsetTxPower
`func (o *InterfaceRequest) UnsetTxPower()`

UnsetTxPower ensures that no value is present for TxPower, not even an explicit nil
### GetUntaggedVlan

`func (o *InterfaceRequest) GetUntaggedVlan() NestedVLANRequest`

GetUntaggedVlan returns the UntaggedVlan field if non-nil, zero value otherwise.

### GetUntaggedVlanOk

`func (o *InterfaceRequest) GetUntaggedVlanOk() (*NestedVLANRequest, bool)`

GetUntaggedVlanOk returns a tuple with the UntaggedVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntaggedVlan

`func (o *InterfaceRequest) SetUntaggedVlan(v NestedVLANRequest)`

SetUntaggedVlan sets UntaggedVlan field to given value.

### HasUntaggedVlan

`func (o *InterfaceRequest) HasUntaggedVlan() bool`

HasUntaggedVlan returns a boolean if a field has been set.

### SetUntaggedVlanNil

`func (o *InterfaceRequest) SetUntaggedVlanNil(b bool)`

 SetUntaggedVlanNil sets the value for UntaggedVlan to be an explicit nil

### UnsetUntaggedVlan
`func (o *InterfaceRequest) UnsetUntaggedVlan()`

UnsetUntaggedVlan ensures that no value is present for UntaggedVlan, not even an explicit nil
### GetTaggedVlans

`func (o *InterfaceRequest) GetTaggedVlans() []int32`

GetTaggedVlans returns the TaggedVlans field if non-nil, zero value otherwise.

### GetTaggedVlansOk

`func (o *InterfaceRequest) GetTaggedVlansOk() (*[]int32, bool)`

GetTaggedVlansOk returns a tuple with the TaggedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaggedVlans

`func (o *InterfaceRequest) SetTaggedVlans(v []int32)`

SetTaggedVlans sets TaggedVlans field to given value.

### HasTaggedVlans

`func (o *InterfaceRequest) HasTaggedVlans() bool`

HasTaggedVlans returns a boolean if a field has been set.

### GetMarkConnected

`func (o *InterfaceRequest) GetMarkConnected() bool`

GetMarkConnected returns the MarkConnected field if non-nil, zero value otherwise.

### GetMarkConnectedOk

`func (o *InterfaceRequest) GetMarkConnectedOk() (*bool, bool)`

GetMarkConnectedOk returns a tuple with the MarkConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkConnected

`func (o *InterfaceRequest) SetMarkConnected(v bool)`

SetMarkConnected sets MarkConnected field to given value.

### HasMarkConnected

`func (o *InterfaceRequest) HasMarkConnected() bool`

HasMarkConnected returns a boolean if a field has been set.

### GetWirelessLans

`func (o *InterfaceRequest) GetWirelessLans() []int32`

GetWirelessLans returns the WirelessLans field if non-nil, zero value otherwise.

### GetWirelessLansOk

`func (o *InterfaceRequest) GetWirelessLansOk() (*[]int32, bool)`

GetWirelessLansOk returns a tuple with the WirelessLans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirelessLans

`func (o *InterfaceRequest) SetWirelessLans(v []int32)`

SetWirelessLans sets WirelessLans field to given value.

### HasWirelessLans

`func (o *InterfaceRequest) HasWirelessLans() bool`

HasWirelessLans returns a boolean if a field has been set.

### GetVrf

`func (o *InterfaceRequest) GetVrf() NestedVRFRequest`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *InterfaceRequest) GetVrfOk() (*NestedVRFRequest, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *InterfaceRequest) SetVrf(v NestedVRFRequest)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *InterfaceRequest) HasVrf() bool`

HasVrf returns a boolean if a field has been set.

### SetVrfNil

`func (o *InterfaceRequest) SetVrfNil(b bool)`

 SetVrfNil sets the value for Vrf to be an explicit nil

### UnsetVrf
`func (o *InterfaceRequest) UnsetVrf()`

UnsetVrf ensures that no value is present for Vrf, not even an explicit nil
### GetTags

`func (o *InterfaceRequest) GetTags() []NestedTagRequest`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InterfaceRequest) GetTagsOk() (*[]NestedTagRequest, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InterfaceRequest) SetTags(v []NestedTagRequest)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InterfaceRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *InterfaceRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *InterfaceRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *InterfaceRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *InterfaceRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


