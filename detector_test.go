package bscexorcist

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

func TestDetectSandwichForBundle(t *testing.T) {
	tests := []struct {
		name    string
		logs    [][]*types.Log
		wantErr bool
	}{
		{
			name:    "testCase1",
			logs:    testCase1,
			wantErr: true,
		},
		{
			name:    "testCase2",
			logs:    testCase2,
			wantErr: true,
		},
		{
			name:    "testCase3",
			logs:    testCase3,
			wantErr: true,
		},
		{
			name:    "testCase4",
			logs:    testCase4,
			wantErr: true,
		},
		{
			name:    "testCase5",
			logs:    testCase5,
			wantErr: true,
		},
		{
			name:    "testCase6DODO",
			logs:    testCase6DODO,
			wantErr: true,
		},
		{
			name:    "testCase0",
			logs:    testCase0,
			wantErr: true,
		},
		{
			name:    "testCase7PancakeV4",
			logs:    testCase7PancakeV4,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DetectSandwichForBundle(tt.logs); (err != nil) != tt.wantErr {
				t.Errorf("DetectSandwichForBundle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

var (
	testCase0 = [][]*types.Log{
		// tx1
		{
			{
				Address: common.HexToAddress("0x5F4D3c0538cAf488053AaE0341f5A4A207F7Ff8e"),
				Topics: []common.Hash{
					common.HexToHash("0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822"),
					common.HexToHash("0x0000000000000000000000002104ba06a100d8cb1a3985237842d5647c157ab8"),
					common.HexToHash("0x0000000000000000000000002104ba06a100d8cb1a3985237842d5647c157ab8"),
				},
				Data: hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001a38ea878535194ec0000000000000000000000000000000000000000000000aad7ffa9942a5d1c880000000000000000000000000000000000000000000000000000000000000000"),
			},
		},
		// tx2
		{
			{
				Address: common.HexToAddress("0x5F4D3c0538cAf488053AaE0341f5A4A207F7Ff8e"),
				Topics: []common.Hash{
					common.HexToHash("0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822"),
					common.HexToHash("0x000000000000000000000000d9c500dff816a1da21a48a732d3498bf09dc9aeb"),
					common.HexToHash("0x000000000000000000000000c7e53484f23333b722d47cc58e93f86806db6220"),
				},
				Data: hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003d1e994fc5cfcdff3e00000000000000000000000000000000000000000000152d02c7e14af68000190000000000000000000000000000000000000000000000000000000000000000"),
			},
		},
		// tx3
		{
			{
				Address: common.HexToAddress("0x5F4D3c0538cAf488053AaE0341f5A4A207F7Ff8e"),
				Topics: []common.Hash{
					common.HexToHash("0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822"),
					common.HexToHash("0x0000000000000000000000002104ba06a100d8cb1a3985237842d5647c157ab8"),
					common.HexToHash("0x0000000000000000000000002104ba06a100d8cb1a3985237842d5647c157ab8"),
				},
				Data: hexutil.MustDecode("0x0000000000000000000000000000000000000000000000aad7ffa9942a5d1c870000000000000000000000000000000000000000000000aad7ffa9942a5d1c8800000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000aca5d502e51879a5bd"),
			},
		},
	}

	testCase1 = [][]*types.Log{
		// tx1
		{
			{
				Address: common.HexToAddress("0xE21d99B42ECCbb7eA0e185f561B0a80fE1FC4d1C"),
				Topics: []common.Hash{
					common.HexToHash("0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822"),
					common.HexToHash("0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822"),
					common.HexToHash("0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822"),
				},
				Data: hexutil.MustDecode("0x00000000000000000000000000000000000000000000001a2241a49c6b6245af000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010d978570f31444e0ce7"),
			},
		},
		// tx2
		{
			{
				Address: common.HexToAddress("0xE21d99B42ECCbb7eA0e185f561B0a80fE1FC4d1C"),
				Topics: []common.Hash{
					common.HexToHash("0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822"),
					common.HexToHash("0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822"),
					common.HexToHash("0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822"),
				},
				Data: hexutil.MustDecode("0x00000000000000000000000000000000000000000000000acf3b1b8894e40000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000006dedd5cc45eb0799505"),
			},
		},
		// tx3
		{
			{
				Address: common.HexToAddress("0xE21d99B42ECCbb7eA0e185f561B0a80fE1FC4d1C"),
				Topics: []common.Hash{
					common.HexToHash("0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822"),
					common.HexToHash("0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822"),
					common.HexToHash("0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822"),
				},
				Data: hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000058f71f3c5037542b775000000000000000000000000000000000000000000000008b621b4b388f9b2590000000000000000000000000000000000000000000000000000000000000000"),
			},
		},
		// tx4
		{
			{
				Address: common.HexToAddress("0xE21d99B42ECCbb7eA0e185f561B0a80fE1FC4d1C"),
				Topics: []common.Hash{
					common.HexToHash("0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822"),
					common.HexToHash("0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822"),
					common.HexToHash("0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822"),
				},
				Data: hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000b4a06634a2dcf0b5572000000000000000000000000000000000000000000000011823e3b2b071b29270000000000000000000000000000000000000000000000000000000000000000"),
			},
		},
	}

	testCase2 = [][]*types.Log{
		// tx1
		{
			{
				Address: common.HexToAddress("0x8BC4BE4d483d1B3EA5Bf7D5a392A3b4bD9A874A9"),
				Topics: []common.Hash{
					common.HexToHash("0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822"),
					common.HexToHash("0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822"),
					common.HexToHash("0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822"),
				},
				Data: hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000b8807aba155303280000000000000000000000000000000000000000002b58249b01a661c0bcd08c0000000000000000000000000000000000000000000000000000000000000000"),
			},
		},
		// tx2
		{
			{
				Address: common.HexToAddress("0x8BC4BE4d483d1B3EA5Bf7D5a392A3b4bD9A874A9"),
				Topics: []common.Hash{
					common.HexToHash("0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822"),
					common.HexToHash("0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822"),
					common.HexToHash("0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822"),
				},
				Data: hexutil.MustDecode("0x00000000000000000000000000000000000000000002fca1f25b3dba8d6764410000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000011f852d27d9ffa52"),
			},
		},
		// tx3
		{
			{
				Address: common.HexToAddress("0x8BC4BE4d483d1B3EA5Bf7D5a392A3b4bD9A874A9"),
				Topics: []common.Hash{
					common.HexToHash("0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822"),
					common.HexToHash("0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822"),
					common.HexToHash("0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822"),
				},
				Data: hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000006de97e09bd1800000000000000000000000000000000000000000000000fb17f35d5ccd60f325f5e0000000000000000000000000000000000000000000000000000000000000000"),
			},
		},
		// tx4
		{
			{
				Address: common.HexToAddress("0x8BC4BE4d483d1B3EA5Bf7D5a392A3b4bD9A874A9"),
				Topics: []common.Hash{
					common.HexToHash("0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822"),
					common.HexToHash("0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822"),
					common.HexToHash("0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822"),
				},
				Data: hexutil.MustDecode("0x0000000000000000000000000000000000000000001156750acd75c0b37eb9d10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000076d82e76d8bdbca2"),
			},
		},
		// tx5
		{
			{
				Address: common.HexToAddress("0x8BC4BE4d483d1B3EA5Bf7D5a392A3b4bD9A874A9"),
				Topics: []common.Hash{
					common.HexToHash("0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822"),
					common.HexToHash("0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822"),
					common.HexToHash("0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822"),
				},
				Data: hexutil.MustDecode("0x0000000000000000000000000000000000000000001a01af903430a10d3e16ba0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000073f8036b326e0c03"),
			},
		},
	}

	testCase3 = [][]*types.Log{
		// tx1
		{
			{
				Address: common.HexToAddress("0x9Ae10A97634Ff4Cff2397D35DDD7f62C2446667D"),
				Topics: []common.Hash{
					common.HexToHash("0x19b47279256b2a23a1665c810c8d55a1758940ee09377d4f8d26497a3577dc83"),
					common.HexToHash("0x19b47279256b2a23a1665c810c8d55a1758940ee09377d4f8d26497a3577dc83"),
					common.HexToHash("0x19b47279256b2a23a1665c810c8d55a1758940ee09377d4f8d26497a3577dc83"),
				},
				Data: hexutil.MustDecode("0xfffffffffffffffffffffffffffffffffffffffffffff32c5190561236a289a00000000000000000000000000000000000000000000000000d133dbb20f49e85000000000000000000000000000000000000000001031a64b244a9110948bbe0000000000000000000000000000000000000000000000a10991d832477b35a73fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe4fb3000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001c4709ca0c20"),
			},
		},
		// tx2
		{
			{
				Address: common.HexToAddress("0x9Ae10A97634Ff4Cff2397D35DDD7f62C2446667D"),
				Topics: []common.Hash{
					common.HexToHash("0x19b47279256b2a23a1665c810c8d55a1758940ee09377d4f8d26497a3577dc83"),
					common.HexToHash("0x19b47279256b2a23a1665c810c8d55a1758940ee09377d4f8d26497a3577dc83"),
					common.HexToHash("0x19b47279256b2a23a1665c810c8d55a1758940ee09377d4f8d26497a3577dc83"),
				},
				Data: hexutil.MustDecode("0xfffffffffffffffffffffffffffffffffffffffffffffec80c99f25d775f4984000000000000000000000000000000000000000000000000013fbe85edc90000000000000000000000000000000000000000000001033a289b1564a5e06e79f1000000000000000000000000000000000000000000000a10991d832477b35a73fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe4fbd0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002b381cb8400"),
			},
		},
		// tx3
		{
			{
				Address: common.HexToAddress("0x9Ae10A97634Ff4Cff2397D35DDD7f62C2446667D"),
				Topics: []common.Hash{
					common.HexToHash("0x19b47279256b2a23a1665c810c8d55a1758940ee09377d4f8d26497a3577dc83"),
					common.HexToHash("0x19b47279256b2a23a1665c810c8d55a1758940ee09377d4f8d26497a3577dc83"),
					common.HexToHash("0x19b47279256b2a23a1665c810c8d55a1758940ee09377d4f8d26497a3577dc83"),
				},
				Data: hexutil.MustDecode("0x0000000000000000000000000000000000000000000007b2357632c1df3813d2fffffffffffffffffffffffffffffffffffffffffffffffff82218dadfb22c48000000000000000000000000000000000000000001027210cec6e24142da238c000000000000000000000000000000000000000000000a10991d832477b35a73fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe4f8100000000000000000000000000000000000000000000000010a4f26ee3c5117f0000000000000000000000000000000000000000000000000000000000000000"),
			},
		},
		// tx4
		{
			{
				Address: common.HexToAddress("0x9Ae10A97634Ff4Cff2397D35DDD7f62C2446667D"),
				Topics: []common.Hash{
					common.HexToHash("0x19b47279256b2a23a1665c810c8d55a1758940ee09377d4f8d26497a3577dc83"),
					common.HexToHash("0x19b47279256b2a23a1665c810c8d55a1758940ee09377d4f8d26497a3577dc83"),
					common.HexToHash("0x19b47279256b2a23a1665c810c8d55a1758940ee09377d4f8d26497a3577dc83"),
				},
				Data: hexutil.MustDecode("0x00000000000000000000000000000000000000000000052178f9772bea25628bfffffffffffffffffffffffffffffffffffffffffffffffffac821a82c62031a00000000000000000000000000000000000000000101ed56dfbb168c2a220ce7000000000000000000000000000000000000000000000a10991d832477b35a73fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe4f580000000000000000000000000000000000000000000000000b18a19f428360ff0000000000000000000000000000000000000000000000000000000000000000"),
			},
		},
	}

	testCase4 = [][]*types.Log{
		// tx1
		{
			{
				Address: common.HexToAddress("0xb6Bb744FB59fa399D09f67Ae3634942F533B577f"),
				Topics: []common.Hash{
					common.HexToHash("0x19b47279256b2a23a1665c810c8d55a1758940ee09377d4f8d26497a3577dc83"),
					common.HexToHash("0x00000000000000000000000000000000c5d17f1bb0245b5eb1cf27c28115996e"),
					common.HexToHash("0x00000000000000000000000000000000c5d17f1bb0245b5eb1cf27c28115996e"),
				},
				Data: hexutil.MustDecode("0xffffffffffffffffffffffffffffffffffffffffffffffff260af37100191b240000000000000000000000000000000000000000000000024b272c4a030060000000000000000000000000000000000000000001a4f8d670605487c25e5848ae0000000000000000000000000000000000000000000001127b950b545353219f00000000000000000000000000000000000000000000000000000000000026dc0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000180cbee2ac519"),
			},
		},
		// tx2
		{
			{
				Address: common.HexToAddress("0xb6Bb744FB59fa399D09f67Ae3634942F533B577f"),
				Topics: []common.Hash{
					common.HexToHash("0x19b47279256b2a23a1665c810c8d55a1758940ee09377d4f8d26497a3577dc83"),
					common.HexToHash("0x0000000000000000000000000f4a1d7fdf4890be35e71f3e0bbc4a0ec377eca3"),
					common.HexToHash("0x0000000000000000000000000f4a1d7fdf4890be35e71f3e0bbc4a0ec377eca3"),
				},
				Data: hexutil.MustDecode("0xffffffffffffffffffffffffffffffffffffffffffffffffe34853336dfc15ef0000000000000000000000000000000000000000000000004db73254763000000000000000000000000000000000000000000001a541501ac1e373df9a4cab350000000000000000000000000000000000000000000001127b950b545353219f00000000000000000000000000000000000000000000000000000000000026e90000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000032ee841b8000"),
			},
		},
		// tx3
		{
			{
				Address: common.HexToAddress("0xb6Bb744FB59fa399D09f67Ae3634942F533B577f"),
				Topics: []common.Hash{
					common.HexToHash("0x19b47279256b2a23a1665c810c8d55a1758940ee09377d4f8d26497a3577dc83"),
					common.HexToHash("0x00000000000000000000000000000000c5d17f1bb0245b5eb1cf27c28115996e"),
					common.HexToHash("0x00000000000000000000000000000000c5d17f1bb0245b5eb1cf27c28115996e"),
				},
				Data: hexutil.MustDecode("0x000000000000000000000000000000000000000000000000d9f50c8effe6e4dcfffffffffffffffffffffffffffffffffffffffffffffffdb4734e6942e114920000000000000000000000000000000000000001a3c40f235663b5146d9e256c00000000000000000000000000000000000000000000033610b8541f16501c5e00000000000000000000000000000000000000000000000000000000000026a200000000000000000000000000000000000000000000000000008ed7277121990000000000000000000000000000000000000000000000000000000000000000"),
			},
		},
	}

	testCase5 = [][]*types.Log{
		// tx1
		{
			{
				Address: common.HexToAddress("0xcF59B8C8BAA2dea520e3D549F97d4e49aDE17057"),
				Topics: []common.Hash{
					common.HexToHash("0x19b47279256b2a23a1665c810c8d55a1758940ee09377d4f8d26497a3577dc83"),
					common.HexToHash("0x000000000000000000000000b5cb0555c0c51e603ead62c6437da65372e4e1b0"),
					common.HexToHash("0x000000000000000000000000b5cb0555c0c51e603ead62c6437da65372e4e1b0"),
				},
				Data: hexutil.MustDecode("0x00000000000000000000000000000000000000000000457023c6a77de0f332d2ffffffffffffffffffffffffffffffffffffffffffffffc9e4ee9e3f80a2abde0000000000000000000000000000000000000000001f3d7fa22fbeaa5c1d89e2000000000000000000000000000000000000000000000006e27bd7319038a46bfffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdaa6b000000000000000000000000000000000000000000000000962c3e7ee2365c120000000000000000000000000000000000000000000000000000000000000000"),
			},
		},
		// tx2
		{
			{
				Address: common.HexToAddress("0x28e2Ea090877bF75740558f6BFB36A5ffeE9e9dF"),
				Topics: []common.Hash{
					common.HexToHash("0x40e9cecb9f5f1f1c5b9c97dec2917b7ee92e57ba5563708daca94dd84ad7112f"),
					common.HexToHash("0xc012e144f83cd4c616704b5391205ad0dc19719ca9f89b739a67a9b1d5316f17"),
					common.HexToHash("0x000000000000000000000000c0fab674ff7ddf8b891495ba9975b0fe1dcac735"),
				},
				Data: hexutil.MustDecode("0xffffffffffffffffffffffffffffffffffffffffffffe76eefee5f2095db0800000000000000000000000000000000000000000000000063030b852e1e7ba6300000000000000000000000000000000000000000201e4fb7bb2dc5c7940e7e13000000000000000000000000000000000000000000392bbff455b2bb9dafab33ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff5dd2000000000000000000000000000000000000000000000000000000000000000a"),
			}, {
				Address: common.HexToAddress("0xcF59B8C8BAA2dea520e3D549F97d4e49aDE17057"),
				Topics: []common.Hash{
					common.HexToHash("0x19b47279256b2a23a1665c810c8d55a1758940ee09377d4f8d26497a3577dc83"),
					common.HexToHash("0x0000000000000000000000007a7ad9aa93cd0a2d0255326e5fb145cec14997ff"),
					common.HexToHash("0x000000000000000000000000927c4306a908362694eece7e4f45c1624dd51721"),
				},
				Data: hexutil.MustDecode("0x000000000000000000000000000000000000000000000a0235f4380f2b3e7800ffffffffffffffffffffffffffffffffffffffffffffffffffdf98888f57ad3a0000000000000000000000000000000000000000001a8895ed39cbb97d40ba1b000000000000000000000000000000000000000000000006e27bd7319038a46bfffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd9da900000000000000000000000000000000000000000000000015a543308721dedd0000000000000000000000000000000000000000000000000000000000000000"),
			}, {
				Address: common.HexToAddress("0x81C7294b66955824BC04acB642ae8dC58e6cE507"),
				Topics: []common.Hash{
					common.HexToHash("0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67"),
					common.HexToHash("0x0000000000000000000000007a7ad9aa93cd0a2d0255326e5fb145cec14997ff"),
					common.HexToHash("0x000000000000000000000000927c4306a908362694eece7e4f45c1624dd51721"),
				},
				Data: hexutil.MustDecode("0x0000000000000000000000000000000000000000000009194852618227501000ffffffffffffffffffffffffffffffffffffffffffffffdb7dc39bf9b68a3d1400000000000000000000000000000000000000001ec15609a9c3c315b47cc4390000000000000000000000000000000000000000000002181e7c4540269d76a3ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff5a6f"),
			}, {
				Address: common.HexToAddress("0x28e2Ea090877bF75740558f6BFB36A5ffeE9e9dF"),
				Topics: []common.Hash{
					common.HexToHash("0x40e9cecb9f5f1f1c5b9c97dec2917b7ee92e57ba5563708daca94dd84ad7112f"),
					common.HexToHash("0xb2008c0126f1adeed67d42da44ac4b5c5abd9ef156bd93c106e082cb10e04552"),
					common.HexToHash("0x000000000000000000000000c0fab674ff7ddf8b891495ba9975b0fe1dcac735"),
				},
				Data: hexutil.MustDecode("0xffffffffffffffffffffffffffffffffffffffffffffff712f87ca2fadd49f5f0000000000000000000000000000000000000000000000023f8adb35e075df3900000000000000000000000000000000000000000000000000000001000276a40000000000000000000000000000000000000000000000000000000000000000fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff276180000000000000000000000000000000000000000000000000000000000000008"),
			}, {
				Address: common.HexToAddress("0x28e2Ea090877bF75740558f6BFB36A5ffeE9e9dF"),
				Topics: []common.Hash{
					common.HexToHash("0x40e9cecb9f5f1f1c5b9c97dec2917b7ee92e57ba5563708daca94dd84ad7112f"),
					common.HexToHash("0xed461e6026d36fa5c289a288cc450fb7556ad2e2642fcda0d173a284140d96bd"),
					common.HexToHash("0x000000000000000000000000c0fab674ff7ddf8b891495ba9975b0fe1dcac735"),
				},
				Data: hexutil.MustDecode("0xffffffffffffffffffffffffffffffffffffffffffffff17125e2972fc1198000000000000000000000000000000000000000000000000005a740cfe34ef868d000000000000000000000000000000000000000000006c756dcfcc1d9835b57e00000000000000000000000000000000000000000000000000592f0d2b1cd66dfffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc5a6c0000000000000000000000000000000000000000000000000000000000000005"),
			},
		},
		// tx3
		{
			{
				Address: common.HexToAddress("0xcF59B8C8BAA2dea520e3D549F97d4e49aDE17057"),
				Topics: []common.Hash{
					common.HexToHash("0x19b47279256b2a23a1665c810c8d55a1758940ee09377d4f8d26497a3577dc83"),
					common.HexToHash("0x000000000000000000000000695a122b5e936f82719217c4823025bff07b4d51"),
					common.HexToHash("0x000000000000000000000000695a122b5e936f82719217c4823025bff07b4d51"),
				},
				Data: hexutil.MustDecode("0xffffffffffffffffffffffffffffffffffffffffffffb746104a3e2f947dfe4900000000000000000000000000000000000000000000001b0d88b0e03faeaa110000000000000000000000000000000000000000201f03bbdfb9eeaa39b2be67000000000000000000000000000000000000000000169e8c5e068dbeecfd0694ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff5dd40000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003a81c4e7fdc86c"),
			},
		},
		// tx4
		{
			{
				Address: common.HexToAddress("0xcF59B8C8BAA2dea520e3D549F97d4e49aDE17057"),
				Topics: []common.Hash{
					common.HexToHash("0x19b47279256b2a23a1665c810c8d55a1758940ee09377d4f8d26497a3577dc83"),
					common.HexToHash("0x000000000000000000000000695a122b5e936f82719217c4823025bff07b4d51"),
					common.HexToHash("0x000000000000000000000000695a122b5e936f82719217c4823025bff07b4d51"),
				},
				Data: hexutil.MustDecode("0xfffffffffffffffffffffffffffffffffffffffffffff949fea54034d018f33900000000000000000000000000000000000000000000001b0d88b0e03faeaa110000000000000000000000000000000000000000201fb01f52e928503fcc55ce0000000000000000000000000000000000000000004dd7131578ce1d99162e57ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff5dd60000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003a81c4e7fdc914"),
			},
		},
	}

	// PancakeSwap V4 sandwich attack test case
	// From online.logs - 3 swaps on the same pool with sandwich pattern
	// Front run: amount0 negative (token0 out), amount1 positive (token1 in) -> token1 to token0 (IsToken0To1 = false)
	// Trigger: same direction as front run
	// Back run: amount0 positive (token0 in), amount1 negative (token1 out) -> token0 to token1 (IsToken0To1 = true)
	testCase7PancakeV4 = [][]*types.Log{
		// tx1 - Front run
		{
			{
				Address: common.HexToAddress("0xa0ffb9c1ce1fe56963b0321b32e7a030211405bb"),
				Topics: []common.Hash{
					common.HexToHash("0x04206ad2b7c0f463bff3dd4f33c5735b0f2957a351e4f79763a4fa9e775dd237"),
					common.HexToHash("0x9a5d9822856a5b78491108a85725d4488352ccd5dcff2f8dbcd0a5a4b5edb3be"),
					common.HexToHash("0x00000000000000000000000048D9cb0C40287A377e961925BD89F777FfDc4eed"),
				},
				Data: hexutil.MustDecode("0xfffffffffffffffffffffffffffffffffffffffffffffff4cc8e67ac48193c2900000000000000000000000000000000000000000000015cddb347054b809f3b00000000000000000000000000000000000000058bd524e66275ce25175b6c74000000000000000000000000000000000000000000001372e02bfd18e3d45a3200000000000000000000000000000000000000000000000000000000000085d800000000000000000000000000000000000000000000000000000000000000630000000000000000000000000000000000000000000000000000000000000020"),
			},
		},
		// tx2 - Trigger (victim)
		{
			{
				Address: common.HexToAddress("0xa0ffb9c1ce1fe56963b0321b32e7a030211405bb"),
				Topics: []common.Hash{
					common.HexToHash("0x04206ad2b7c0f463bff3dd4f33c5735b0f2957a351e4f79763a4fa9e775dd237"),
					common.HexToHash("0x9a5d9822856a5b78491108a85725d4488352ccd5dcff2f8dbcd0a5a4b5edb3be"),
					common.HexToHash("0x000000000000000000000000000000000000000000000000000000000000dead"), // PancakeSwapInfinityHookAdapter placeholder
				},
				Data: hexutil.MustDecode("0xffffffffffffffffffffffffffffffffffffffffffffffef4b1c01c1c845e1b90000000000000000000000000000000000000000000001f87765aaa21c6b2110000000000000000000000000000000000000000571e4f2c1e243270f9bea3862000000000000000000000000000000000000000000001372e02bfd18e3d45a32000000000000000000000000000000000000000000000000000000000000846700000000000000000000000000000000000000000000000000000000000000630000000000000000000000000000000000000000000000000000000000000020"),
			},
		},
		// tx3 - Back run
		{
			{
				Address: common.HexToAddress("0xa0ffb9c1ce1fe56963b0321b32e7a030211405bb"),
				Topics: []common.Hash{
					common.HexToHash("0x04206ad2b7c0f463bff3dd4f33c5735b0f2957a351e4f79763a4fa9e775dd237"),
					common.HexToHash("0x9a5d9822856a5b78491108a85725d4488352ccd5dcff2f8dbcd0a5a4b5edb3be"),
					common.HexToHash("0x00000000000000000000000048D9cb0C40287A377e961925BD89F777FfDc4eed"),
				},
				Data: hexutil.MustDecode("0x00000000000000000000000000000000000000000000000b9de0e60d11af2732fffffffffffffffffffffffffffffffffffffffffffffea3224cb8fab47f60c6000000000000000000000000000000000000000583d48d17745ab37c7ce460ee000000000000000000000000000000000000000000001372e02bfd18e3d45a32000000000000000000000000000000000000000000000000000000000000856700000000000000000000000000000000000000000000000000000000000000630000000000000000000000000000000000000000000000000000000000000020"),
			},
		},
	}

	testCase6DODO = [][]*types.Log{
		// tx1
		{
			{
				Address: common.HexToAddress("0x47423Dbe87f337A4a5c17Cc57d09B971d501DD14"),
				Topics: []common.Hash{
					common.HexToHash("0xc2c0245e056d5fb095f04cd6373bc770802ebd1e6c918eb78fdef843cdb37b0f"),
				},
				Data: hexutil.MustDecode("0x0000000000000000000000008ac76a51cc950d9822d68b83fe1ad97b32cd580d00000000000000000000000055d398326f99059ff775485246999027b31979550000000000000000000000000000000000000000000005884678fede0dd416b10000000000000000000000000000000000000000000005875c290aca3bfc4a7f00000000000000000000000042a7898c2cbfd351603551ff626c6493e4fef75100000000000000000000000042a7898c2cbfd351603551ff626c6493e4fef751"),
			},
		},
		// tx2
		{
			{
				Address: common.HexToAddress("0x47423Dbe87f337A4a5c17Cc57d09B971d501DD14"),
				Topics: []common.Hash{
					common.HexToHash("0xc2c0245e056d5fb095f04cd6373bc770802ebd1e6c918eb78fdef843cdb37b0f"),
				},
				Data: hexutil.MustDecode("0x0000000000000000000000008ac76a51cc950d9822d68b83fe1ad97b32cd580d00000000000000000000000055d398326f99059ff775485246999027b31979550000000000000000000000000000000000000000000001e239086fb266dbd0000000000000000000000000000000000000000000000001dfc3e5c4d000738b970000000000000000000000004a99693dbc7545efc4b39431f163adcb4ed541fb0000000000000000000000004a99693dbc7545efc4b39431f163adcb4ed541fb"),
			},
		},
		// tx3
		{
			{
				Address: common.HexToAddress("0x47423Dbe87f337A4a5c17Cc57d09B971d501DD14"),
				Topics: []common.Hash{
					common.HexToHash("0xc2c0245e056d5fb095f04cd6373bc770802ebd1e6c918eb78fdef843cdb37b0f"),
				},
				Data: hexutil.MustDecode("0x00000000000000000000000055d398326f99059ff775485246999027b31979550000000000000000000000008ac76a51cc950d9822d68b83fe1ad97b32cd580d0000000000000000000000000000000000000000000005875c290aca3bfc4a7f00000000000000000000000000000000000000000000058a746dece7b6fba04a000000000000000000000000946d7590da72429897a765972ad7922e7040483d000000000000000000000000946d7590da72429897a765972ad7922e7040483d"),
			},
		},
	}
)

// --- Regression: a single multi-hop transaction must not fabricate a sandwich ---
//
// PancakeV3 swap signature; reused to build synthetic V3 swap logs.
const v3SwapSig = "0x19b47279256b2a23a1665c810c8d55a1758940ee09377d4f8d26497a3577dc83"

// int256Bytes encodes a signed int into a 32-byte big-endian two's-complement word,
// matching tools.DecodeSignedInt256.
func int256Bytes(v int64) []byte {
	b := big.NewInt(v)
	if v < 0 {
		b.Add(new(big.Int).Lsh(big.NewInt(1), 256), b)
	}
	out := make([]byte, 32)
	b.FillBytes(out)
	return out
}

// v3SwapLog builds a minimal PancakeV3 Swap log. amount0>0/amount1<0 => token0->token1
// (IsToken0To1==true, "sell token0"); amount0<0/amount1>0 => token1->token0 ("buy token0").
func v3SwapLog(pool common.Address, amount0, amount1 int64) *types.Log {
	data := make([]byte, 160) // amount0, amount1, sqrtPriceX96, liquidity, tick
	copy(data[0:32], int256Bytes(amount0))
	copy(data[32:64], int256Bytes(amount1))
	return &types.Log{
		Address: pool,
		Topics:  []common.Hash{common.HexToHash(v3SwapSig)},
		Data:    data,
	}
}

// A honest backrun: one victim tx whose internal multi-hop routes through the same
// pool several times in one direction, followed by independent backrun txs in the
// opposite direction. The per-tx swaps of the victim must collapse to a single net
// leg, so the bundle is [T, F, F, F, F] — not a sandwich.
func TestDetectSandwichForBundle_BackrunNotFlagged(t *testing.T) {
	pool := common.HexToAddress("0x194fb07E51EEb6FeC9Eae4edad6523F8dFF2C7eD")

	victimTx := []*types.Log{
		v3SwapLog(pool, 100, -90),
		v3SwapLog(pool, 100, -90),
		v3SwapLog(pool, 100, -90),
		v3SwapLog(pool, 100, -90),
	}
	backrun := func() []*types.Log { return []*types.Log{v3SwapLog(pool, -100, 90)} }

	bundle := [][]*types.Log{victimTx, backrun(), backrun(), backrun(), backrun()}

	if err := DetectSandwichForBundle(bundle); err != nil {
		t.Fatalf("honest backrun must not be flagged as sandwich, got: %v", err)
	}
}

// A classic sandwich keeps front-run and back-run in separate transactions around
// the victim, so per-tx collapse leaves [T, T, F] intact and it is still detected.
func TestDetectSandwichForBundle_ClassicStillDetected(t *testing.T) {
	pool := common.HexToAddress("0x194fb07E51EEb6FeC9Eae4edad6523F8dFF2C7eD")

	frontRun := []*types.Log{v3SwapLog(pool, 100, -90)} // buy (T)
	victim := []*types.Log{v3SwapLog(pool, 100, -90)}   // buy (T)
	backRun := []*types.Log{v3SwapLog(pool, -100, 90)}  // sell (F)

	bundle := [][]*types.Log{frontRun, victim, backRun}

	if err := DetectSandwichForBundle(bundle); err == nil {
		t.Fatal("classic 3-tx sandwich must still be detected")
	}
}

// --- Regression: a directional Sell-Sell-Buy that fails asset conservation is not a sandwich ---

// signed256 encodes a signed big.Int as a 32-byte big-endian two's-complement word.
func signed256(v *big.Int) []byte {
	b := new(big.Int).Set(v)
	if b.Sign() < 0 {
		b.Add(new(big.Int).Lsh(big.NewInt(1), 256), b)
	}
	out := make([]byte, 32)
	b.FillBytes(out)
	return out
}

// pancakeV4Log builds a PancakeV4 Swap log. amount0>0 => token0->token1 (IsToken0To1==true).
func pancakeV4Log(poolID common.Hash, amount0, amount1 *big.Int) *types.Log {
	data := make([]byte, 224)
	copy(data[0:32], signed256(amount0))
	copy(data[32:64], signed256(amount1))
	return &types.Log{
		Topics: []common.Hash{
			common.HexToHash("0x04206ad2b7c0f463bff3dd4f33c5735b0f2957a351e4f79763a4fa9e775dd237"),
			poolID,
			common.HexToHash("0x0000000000000000000000005f5e8660d388db53a57b0650ab1565d95a21297f"),
		},
		Data: data,
	}
}

// Reproduces production bundle 0x9e62...4062 on PancakeV4 pool 0xf6170...: searcher A's
// incidental sell, then searcher B's multi-leg arbitrage (sell then buy back). The
// directions are Sell-Sell-Buy, but the back-run buys 11.11e18 token0 while the only
// earlier sell is A's 7.09e18 — they do not conserve, so no victim is bracketed and this
// must not be flagged. (Was a false positive under direction-only detection.)
func TestDetectSandwichForBundle_CrossTraderArbNotFlagged(t *testing.T) {
	pool := common.HexToHash("0xf6170c9b48765c56dae252b10e87e9ac51330846000000000000000000000000")
	bn := func(s string) *big.Int { v, _ := new(big.Int).SetString(s, 10); return v }
	bundle := [][]*types.Log{
		{pancakeV4Log(pool, bn("-7094659767506153562"), bn("26428545600792617964"))},  // A: sell token1->token0
		{pancakeV4Log(pool, bn("-15570669439182987264"), bn("57998528022014499813"))}, // B: sell token1->token0
		{pancakeV4Log(pool, bn("11110624239951633190"), bn("-41393089742295897788"))}, // B: buy token0->token1
	}
	if err := DetectSandwichForBundle(bundle); err != nil {
		t.Fatalf("cross-trader arbitrage (no conserving bracket) must not be flagged, got: %v", err)
	}
}

// A clean PancakeV4 sandwich whose back-run is split across two transactions: front-run
// buys ~100 token0-worth, two back-runs sell ~50 each. The pairwise front/back never
// matches, but the conservation check sums the split back-run and still detects it.
func TestDetectSandwichForBundle_SplitBackrunStillDetected(t *testing.T) {
	pool := common.HexToHash("0xabcdef0000000000000000000000000000000000000000000000000000000000")
	bn := func(s string) *big.Int { v, _ := new(big.Int).SetString(s, 10); return v }
	bundle := [][]*types.Log{
		// front-run: buy token1 (amount0>0), acquires 1000 token1
		{pancakeV4Log(pool, bn("500000000000000000000"), bn("-1000000000000000000000"))},
		// victim: also buys token1
		{pancakeV4Log(pool, bn("250000000000000000000"), bn("-480000000000000000000"))},
		// back-run split #1: sell token1 (amount0<0), disposes 500 token1
		{pancakeV4Log(pool, bn("-260000000000000000000"), bn("500000000000000000000"))},
		// back-run split #2: sell token1, disposes 500 token1 -> 500+500 == front 1000
		{pancakeV4Log(pool, bn("-260000000000000000000"), bn("500000000000000000000"))},
	}
	if err := DetectSandwichForBundle(bundle); err == nil {
		t.Fatal("sandwich with a split back-run must still be detected")
	}
}

// A transaction whose majority swap direction is token0->token1 but whose NET token1
// change is negative (it really sold token1) must not be counted as a front-run that
// "bought" token1. tx1 below folds to direction T (two tiny buys outvote one big sell)
// yet nets dT1 = -96; if the sign were dropped via Abs it would falsely conserve against
// tx3's sell of 96 token1 and be flagged.
func TestDetectSandwichForBundle_FoldedDirectionMismatchNotFlagged(t *testing.T) {
	pool := common.HexToHash("0x1111111111111111111111111111111111111111000000000000000000000000")
	bundle := [][]*types.Log{
		{ // tx1: folds to T (2 buys vs 1 sell) but net sells token1 (2+2-100 = -96)
			pancakeV4Log(pool, big.NewInt(10), big.NewInt(-2)),
			pancakeV4Log(pool, big.NewInt(10), big.NewInt(-2)),
			pancakeV4Log(pool, big.NewInt(-50), big.NewInt(100)),
		},
		{pancakeV4Log(pool, big.NewInt(30), big.NewInt(-50))}, // tx2: clean buy (would-be victim)
		{pancakeV4Log(pool, big.NewInt(-48), big.NewInt(96))}, // tx3: sells 96 token1
	}
	if err := DetectSandwichForBundle(bundle); err != nil {
		t.Fatalf("a leg whose net asset direction contradicts its folded direction must not conserve, got: %v", err)
	}
}

// The same signed-flow constraint must apply to the victim. A sandwiched victim is a
// buyer that ends up holding more of the asset; a middle leg that folds to direction T
// but net-sells the asset is not being sandwiched. Here front buys 96 token1, the middle
// leg folds to T but nets dT1 = -96, and back sells 96 token1 — it must not be flagged.
func TestDetectSandwichForBundle_FoldedVictimMismatchNotFlagged(t *testing.T) {
	pool := common.HexToHash("0x2222222222222222222222222222222222222222000000000000000000000000")
	bundle := [][]*types.Log{
		{pancakeV4Log(pool, big.NewInt(48), big.NewInt(-96))}, // front: clean buy of 96 token1
		{ // middle: folds to T (2 buys vs 1 sell) but net sells token1 (2+2-100 = -96)
			pancakeV4Log(pool, big.NewInt(10), big.NewInt(-2)),
			pancakeV4Log(pool, big.NewInt(10), big.NewInt(-2)),
			pancakeV4Log(pool, big.NewInt(-50), big.NewInt(100)),
		},
		{pancakeV4Log(pool, big.NewInt(-48), big.NewInt(96))}, // back: sells 96 token1
	}
	if err := DetectSandwichForBundle(bundle); err != nil {
		t.Fatalf("a victim whose net asset direction contradicts its folded direction must not conserve, got: %v", err)
	}
}

// And to the back-run: a leg that folds to "sell" but net-bought the asset has not
// disposed of anything, so it cannot be the back-run that unwinds a sandwich. Here the
// last leg folds to F (2 sells vs 1 buy) but nets dT1 = +96; matching front (96) + victim
// must not conserve against it.
func TestDetectSandwichForBundle_FoldedBackrunMismatchNotFlagged(t *testing.T) {
	pool := common.HexToHash("0x3333333333333333333333333333333333333333000000000000000000000000")
	bundle := [][]*types.Log{
		{pancakeV4Log(pool, big.NewInt(48), big.NewInt(-96))}, // front: clean buy of 96 token1
		{pancakeV4Log(pool, big.NewInt(25), big.NewInt(-50))}, // victim: clean buy of 50 token1
		{ // back: folds to F (2 sells vs 1 buy) but net buys token1 (-2-2+100 = +96)
			pancakeV4Log(pool, big.NewInt(-1), big.NewInt(2)),
			pancakeV4Log(pool, big.NewInt(-1), big.NewInt(2)),
			pancakeV4Log(pool, big.NewInt(50), big.NewInt(-100)),
		},
	}
	if err := DetectSandwichForBundle(bundle); err != nil {
		t.Fatalf("a back-run whose net asset direction contradicts its folded direction must not conserve, got: %v", err)
	}
}

// A real sandwich preceded by an unrelated same-direction trade must still be detected.
// An incidental third-party buy (50) sits before the attacker's front-run (buy 100); the
// attacker then unwinds exactly 100 in the back-run around the victim. Conservation must
// match the front-run nearest the victim (100) against the back-run (100) and ignore the
// incidental earlier buy — summing every same-direction leg before the victim (150) would
// miss this and hand the attacker a one-swap evasion.
func TestDetectSandwichForBundle_LeadingIncidentalFrontStillDetected(t *testing.T) {
	pool := common.HexToHash("0xaaaa111111111111111111111111111111111111000000000000000000000000")
	bundle := [][]*types.Log{
		{pancakeV4Log(pool, big.NewInt(25), big.NewInt(-50))},  // incidental third-party buy of 50 token1
		{pancakeV4Log(pool, big.NewInt(50), big.NewInt(-100))}, // attacker front-run: buy 100 token1
		{pancakeV4Log(pool, big.NewInt(20), big.NewInt(-40))},  // victim: buy 40 token1
		{pancakeV4Log(pool, big.NewInt(-50), big.NewInt(100))}, // attacker back-run: sell 100 token1
	}
	if err := DetectSandwichForBundle(bundle); err == nil {
		t.Fatal("a sandwich preceded by an incidental same-direction trade must still be detected")
	}
}

// The symmetric case: a real sandwich followed by an unrelated same-direction trade must
// still be detected. The attacker unwinds exactly 100 (back-run) right after the victim;
// an incidental later sell (50) must not inflate the back-run total past tolerance.
func TestDetectSandwichForBundle_TrailingIncidentalBackStillDetected(t *testing.T) {
	pool := common.HexToHash("0xbbbb222222222222222222222222222222222222000000000000000000000000")
	bundle := [][]*types.Log{
		{pancakeV4Log(pool, big.NewInt(50), big.NewInt(-100))}, // attacker front-run: buy 100 token1
		{pancakeV4Log(pool, big.NewInt(20), big.NewInt(-40))},  // victim: buy 40 token1
		{pancakeV4Log(pool, big.NewInt(-50), big.NewInt(100))}, // attacker back-run: sell 100 token1
		{pancakeV4Log(pool, big.NewInt(-25), big.NewInt(50))},  // incidental third-party sell of 50 token1
	}
	if err := DetectSandwichForBundle(bundle); err == nil {
		t.Fatal("a sandwich followed by an incidental same-direction trade must still be detected")
	}
}

// An opposite-direction trade between two same-direction legs breaks the front-run: the
// two buys (40, 60) are separated by a sell (10) and must NOT be stitched into a 100 that
// conserves against the back-run (sell 100). Each is its own front-run candidate (40, 60),
// neither matches 100, so this must not be flagged.
func TestDetectSandwichForBundle_InterruptedFrontNotStitched(t *testing.T) {
	pool := common.HexToHash("0xcccc333333333333333333333333333333333333000000000000000000000000")
	bundle := [][]*types.Log{
		{pancakeV4Log(pool, big.NewInt(20), big.NewInt(-40))},  // buy 40 token1
		{pancakeV4Log(pool, big.NewInt(-5), big.NewInt(10))},   // sell 10 token1 (interrupts the run)
		{pancakeV4Log(pool, big.NewInt(30), big.NewInt(-60))},  // buy 60 token1
		{pancakeV4Log(pool, big.NewInt(20), big.NewInt(-40))},  // victim: buy 40 token1
		{pancakeV4Log(pool, big.NewInt(-50), big.NewInt(100))}, // sell 100 token1
	}
	if err := DetectSandwichForBundle(bundle); err != nil {
		t.Fatalf("front legs split by an opposite-direction trade must not be stitched, got: %v", err)
	}
}

// The symmetric case on the back side: a back-run interrupted by an opposite-direction
// trade must not be stitched. The sells (60, 40) are separated by a buy (10) and must not
// be summed into a 100 that conserves against the front-run (buy 100).
func TestDetectSandwichForBundle_InterruptedBackNotStitched(t *testing.T) {
	pool := common.HexToHash("0xdddd444444444444444444444444444444444444000000000000000000000000")
	bundle := [][]*types.Log{
		{pancakeV4Log(pool, big.NewInt(50), big.NewInt(-100))}, // buy 100 token1
		{pancakeV4Log(pool, big.NewInt(15), big.NewInt(-30))},  // victim: buy 30 token1
		{pancakeV4Log(pool, big.NewInt(-30), big.NewInt(60))},  // sell 60 token1
		{pancakeV4Log(pool, big.NewInt(5), big.NewInt(-10))},   // buy 10 token1 (interrupts the run)
		{pancakeV4Log(pool, big.NewInt(-20), big.NewInt(40))},  // sell 40 token1
	}
	if err := DetectSandwichForBundle(bundle); err != nil {
		t.Fatalf("back legs split by an opposite-direction trade must not be stitched, got: %v", err)
	}
}

// A many-legged reliable-amount pool must still require conservation rather than reverting
// to direction-only flagging. 17 legs shaped Buy-...-Sell where the front-runs buy 1e6
// token1 each and the lone sell disposes 1 token1 do not conserve and must not be flagged.
func TestDetectSandwichForBundle_OverCapNonConservingNotFlagged(t *testing.T) {
	pool := common.HexToHash("0x4444444444444444444444444444444444444444000000000000000000000000")
	oneM := big.NewInt(1_000_000)
	var bundle [][]*types.Log
	for i := 0; i < 16; i++ { // 16 buys of 1e6 token1 (direction T)
		bundle = append(bundle, []*types.Log{pancakeV4Log(pool, big.NewInt(500_000), new(big.Int).Neg(oneM))})
	}
	// one sell of just 1 token1 (direction F) -> 17 legs, has the T,T,F shape but no conservation
	bundle = append(bundle, []*types.Log{pancakeV4Log(pool, big.NewInt(-1), big.NewInt(1))})
	if err := DetectSandwichForBundle(bundle); err != nil {
		t.Fatalf("many-legged pool with the sandwich shape but no conservation must not be flagged, got: %v", err)
	}
}

// On a many-legged pool a genuinely conserving sandwich must still be detected: front buys
// 1000 token1, many victims buy, and the back-run sells 1000 token1 back. 17 legs.
func TestDetectSandwichForBundle_OverCapConservingStillDetected(t *testing.T) {
	pool := common.HexToHash("0x5555555555555555555555555555555555555555000000000000000000000000")
	var bundle [][]*types.Log
	bundle = append(bundle, []*types.Log{pancakeV4Log(pool, big.NewInt(500), big.NewInt(-1000))}) // front: buy 1000 token1
	for i := 0; i < 15; i++ {                                                                     // 15 victim buys
		bundle = append(bundle, []*types.Log{pancakeV4Log(pool, big.NewInt(20), big.NewInt(-40))})
	}
	bundle = append(bundle, []*types.Log{pancakeV4Log(pool, big.NewInt(-480), big.NewInt(1000))}) // back: sell 1000 token1
	if err := DetectSandwichForBundle(bundle); err == nil {
		t.Fatal("many-legged conserving sandwich must still be detected")
	}
}

// A sandwich whose FRONT-RUN is split across two transactions must still be detected:
// front buys 500 + 500 token1, victim buys, back sells 1000 token1; filler buys after the
// victim pad the pool. No single front leg equals the back-run, so only summing the two
// adjacent front legs into one run catches it.
func TestDetectSandwichForBundle_OverCapSplitFrontStillDetected(t *testing.T) {
	pool := common.HexToHash("0x6666666666666666666666666666666666666666000000000000000000000000")
	var bundle [][]*types.Log
	bundle = append(bundle, []*types.Log{pancakeV4Log(pool, big.NewInt(250), big.NewInt(-500))})  // front #1: buy 500 token1
	bundle = append(bundle, []*types.Log{pancakeV4Log(pool, big.NewInt(250), big.NewInt(-500))})  // front #2: buy 500 token1
	bundle = append(bundle, []*types.Log{pancakeV4Log(pool, big.NewInt(20), big.NewInt(-40))})    // victim: buy token1
	bundle = append(bundle, []*types.Log{pancakeV4Log(pool, big.NewInt(-480), big.NewInt(1000))}) // back: sell 1000 token1
	for i := 0; i < 14; i++ {                                                                     // filler buys after victim -> 18 legs
		bundle = append(bundle, []*types.Log{pancakeV4Log(pool, big.NewInt(5), big.NewInt(-10))})
	}
	if err := DetectSandwichForBundle(bundle); err == nil {
		t.Fatal("many-legged sandwich with a split front-run must still be detected")
	}
}

// A pool with the sandwich shape but no conservation must not be flagged just because it
// has many legs. 64 large buys of token1 plus one sell of 1 token1 do not conserve and
// must not be flagged.
func TestDetectSandwichForBundle_ManyLegsNonConservingNotFlagged(t *testing.T) {
	pool := common.HexToHash("0x7777777777777777777777777777777777777777000000000000000000000000")
	var bundle [][]*types.Log
	for i := 0; i < 64; i++ {
		bundle = append(bundle, []*types.Log{pancakeV4Log(pool, big.NewInt(500_000), big.NewInt(-1_000_000))})
	}
	bundle = append(bundle, []*types.Log{pancakeV4Log(pool, big.NewInt(-1), big.NewInt(1))})
	if err := DetectSandwichForBundle(bundle); err != nil {
		t.Fatalf("65-leg non-conserving pool must not be flagged, got: %v", err)
	}
}

// A very large non-conserving pool must not be flagged: the conservation totals never
// match, regardless of how many legs share the sandwich shape.
func TestDetectSandwichForBundle_HardCapNonConservingNotFlagged(t *testing.T) {
	pool := common.HexToHash("0x8888888888888888888888888888888888888888000000000000000000000000")
	var bundle [][]*types.Log
	for i := 0; i < 520; i++ { // many same-direction buy legs
		bundle = append(bundle, []*types.Log{pancakeV4Log(pool, big.NewInt(500_000), big.NewInt(-1_000_000))})
	}
	bundle = append(bundle, []*types.Log{pancakeV4Log(pool, big.NewInt(-1), big.NewInt(1))})
	if err := DetectSandwichForBundle(bundle); err != nil {
		t.Fatalf("very large non-conserving pool must not be flagged, got: %v", err)
	}
}

// A genuinely conserving sandwich on a large pool must still be detected, so padding a
// sandwich with many filler trades cannot hide it: front buys 1000 token1, ~500 filler
// buys, back sells 1000 token1 back.
func TestDetectSandwichForBundle_LargeConservingStillDetected(t *testing.T) {
	pool := common.HexToHash("0x9999999999999999999999999999999999999999000000000000000000000000")
	var bundle [][]*types.Log
	bundle = append(bundle, []*types.Log{pancakeV4Log(pool, big.NewInt(500), big.NewInt(-1000))}) // front: buy 1000 token1
	for i := 0; i < 500; i++ {                                                                    // filler buys (502 legs total)
		bundle = append(bundle, []*types.Log{pancakeV4Log(pool, big.NewInt(5), big.NewInt(-10))})
	}
	bundle = append(bundle, []*types.Log{pancakeV4Log(pool, big.NewInt(-480), big.NewInt(1000))}) // back: sell 1000 token1
	if err := DetectSandwichForBundle(bundle); err == nil {
		t.Fatal("a conserving sandwich padded with filler trades must still be detected")
	}
}

// --- Regression: collapse must rely only on direction, not per-protocol amounts ---

const (
	fourMemeBuySig  = "0x7db52723a3b2cdd6164364b3b766e65e540d7be48ffa89582956d8eaebe62942"
	fourMemeSellSig = "0x0a5575b3648bae2210cee56bf33254cc1ddfbc7bf637c0af2ac18b14fb1bae19"
)

// fourMemeSwapLog builds a FourMeme swap log. FourMeme's AmountIn/AmountOut both
// return 0, so any amount-based net direction would silently drop these swaps.
func fourMemeSwapLog(token common.Address, buy bool) *types.Log {
	sig := fourMemeSellSig
	if buy {
		sig = fourMemeBuySig
	}
	data := make([]byte, 32)
	copy(data[12:32], token.Bytes()) // address occupies the low 20 bytes
	return &types.Log{
		Topics: []common.Hash{common.HexToHash(sig)},
		Data:   data,
	}
}

// FourMeme reports direction via IsToken0To1 only; a buy-buy-sell across three
// separate transactions must still be detected.
func TestDetectSandwichForBundle_FourMemeStillDetected(t *testing.T) {
	token := common.HexToAddress("0x000000000000000000000000000000000000dEaD")
	bundle := [][]*types.Log{
		{fourMemeSwapLog(token, true)},  // buy  (T)
		{fourMemeSwapLog(token, true)},  // buy  (T)
		{fourMemeSwapLog(token, false)}, // sell (F)
	}
	if err := DetectSandwichForBundle(bundle); err == nil {
		t.Fatal("FourMeme buy-buy-sell sandwich must still be detected")
	}
}

// dodoSwapLog builds a DODO swap log. tokenFrom<tokenTo => token0->token1 (T).
func dodoSwapLog(from, to common.Address, amountFrom, amountTo int64) *types.Log {
	data := make([]byte, 128)
	copy(data[12:32], from.Bytes())
	copy(data[44:64], to.Bytes())
	big.NewInt(amountFrom).FillBytes(data[64:96])
	big.NewInt(amountTo).FillBytes(data[96:128])
	return &types.Log{
		Topics: []common.Hash{common.HexToHash("0xc2c0245e056d5fb095f04cd6373bc770802ebd1e6c918eb78fdef843cdb37b0f")},
		Data:   data,
	}
}

// DODO's AmountIn/AmountOut are indexed by token0/token1, not input/output, so an
// amount-based net misattributes a reverse swap's token1 amount as token0 output.
// A single transaction that round-trips the same DODO pool must net to no signal,
// leaving [T, T] which is not a sandwich.
func TestDetectSandwichForBundle_DodoRoundTripNotFlagged(t *testing.T) {
	a := common.HexToAddress("0x0000000000000000000000000000000000000001")
	b := common.HexToAddress("0x0000000000000000000000000000000000000002")
	roundTrip := []*types.Log{
		dodoSwapLog(a, b, 1000, 5), // T (token0 -> token1)
		dodoSwapLog(b, a, 5, 1000), // F (token1 -> token0), nets to a tie
	}
	buy := func() []*types.Log { return []*types.Log{dodoSwapLog(a, b, 10, 1)} } // T

	bundle := [][]*types.Log{roundTrip, buy(), buy()}
	if err := DetectSandwichForBundle(bundle); err != nil {
		t.Fatalf("DODO round-trip + two buys [_,T,T] must not be flagged, got: %v", err)
	}
}

// hasDirectionalPattern is the direction-only fallback for unreliable pools (DODO,
// FourMeme); this pins its Buy-Buy-Sell (T,T,F) / Sell-Sell-Buy (F,F,T) detection.
func TestHasDirectionalPattern(t *testing.T) {
	T, F := true, false
	cases := []struct {
		name string
		dirs []bool
		want bool
	}{
		{"buy-buy-sell", []bool{T, T, F}, true},
		{"sell-sell-buy", []bool{F, F, T}, true},
		{"too-short", []bool{T, F}, false},
		{"buy-sell-buy", []bool{T, F, T}, false},
		{"all-buys", []bool{T, T, T}, false},
		{"all-sells", []bool{F, F, F}, false},
		{"backrun-shape", []bool{T, F, F, F, F}, false},
		{"interrupted-ssb", []bool{F, T, F, T, T}, true}, // testCase2 directions
		{"pattern-deep-in", []bool{F, F, F, F, T}, true},
		{"empty", nil, false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := hasDirectionalPattern(c.dirs); got != c.want {
				t.Errorf("hasDirectionalPattern(%v) = %v, want %v", c.dirs, got, c.want)
			}
		})
	}
}
