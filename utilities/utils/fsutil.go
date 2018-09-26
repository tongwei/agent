package utils

import (
	"github.com/shirou/gopsutil/disk"
	"path/filepath"
	"github.com/jaypipes/ghw"
	"fmt"
)

func GetMountpoint(path string) (string, error){
	partitions, err := GetDiskPartitions(false)
	if err != nil {
		return "", err
	}
	path, _ = filepath.EvalSymlinks(path)
	path, _ = filepath.Abs(filepath.Clean(path))
	for {
		if path == "/" {
			break
		}
		_, ok := partitions[path]
		if ok {
			return "", nil
		}
		path = filepath.Dir(path)
	}
	return path, nil
}

func GetDiskPartitions(all bool) (map[string]disk.PartitionStat, error){
	partitions, err := disk.Partitions(all)
	byMountpoint := make(map[string]disk.PartitionStat)
	if err != nil {
		return byMountpoint, err
	}
	for _, partition := range partitions {
		byMountpoint[partition.Mountpoint] = partition
	}
	fmt.Printf("%v\n", byMountpoint)
	return byMountpoint, nil
}

func GetDevice(path string) (string, error) {
	partitions, err := GetDiskPartitions(false)
	if err != nil {
		return "", err
	}
	p, err := GetMountpoint(path)
        if err != nil {
                return "", err
        }
	v, _ := partitions[p]
	return v.Device, nil
}

func GetRawDevice(path string) (string, error) {
	device, err := GetDevice(path)
	if err != nil {
		return "", err
	}

	block, err := ghw.Block()
	if err != nil {
		return "", err
	}
	name := ""
	fmt.Printf("%v\n", device)
	for _, rawdisk := range block.Disks {
		fmt.Printf("%v\n", rawdisk)
		for _, part := range rawdisk.Partitions {
			fmt.Printf("%v\n", part)
			if fmt.Sprintf("/dev/%s", part.Name) == device {
				name = fmt.Sprintf("/dev/%s", rawdisk.Name)
			}
		}
	}
	return name, nil
}
