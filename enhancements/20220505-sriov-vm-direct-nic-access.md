# SR-IOV (Single root I/O Virtualization)

SR-IOV enables VMs to directly access the NIC without passing data through the virtualization layer. 
It also allows multiple VMs to use the same physical NIC by creating virtual functions (VFs) which 
the VM can configure and own. The SR-IOV driver creates new `eth0v2` and `eth0v3` devices that all 
funnel to the same physical NIC, `eth0`.

By default, this will be turned off, but the way to enable it is by setting the Harvester setting the `sriov-vfs-count` setting to a nonzero value.