#!/bin/bash
set -e

# Create loop devices for ZFS testing
for i in {0..3}; do
    dd if=/dev/zero of=/tmp/zfs_test_$i.img bs=100M count=1
    losetup -f /tmp/zfs_test_$i.img
done

# Get the loop devices we just created
LOOPS=$(losetup -n -O name | tail -n 4)
read -r LOOP1 LOOP2 LOOP3 LOOP4 <<< "$LOOPS"

# Create a mirrored pool
zpool create testpool mirror $LOOP1 $LOOP2

# Create another pool with a single disk
zpool create datapool $LOOP3

# Create some datasets and nested datasets
zfs create testpool/dataset1
zfs create testpool/dataset1/nested1
zfs create testpool/dataset2

# Create a dataset with some properties
zfs create datapool/backup
zfs set compression=on datapool/backup
zfs set quota=50M datapool/backup

# Create some test snapshots
zfs snapshot testpool/dataset1@snap1
zfs snapshot testpool/dataset1@snap2

# Set some properties
zfs set compression=lz4 testpool/dataset1
zfs set quota=25M testpool/dataset2

# Create a small file in one of the datasets
mkdir -p /testpool/dataset1
echo "Test data" > /testpool/dataset1/testfile.txt