package nodebootstrap

// This file was generated by maxpods_generate.go; DO NOT EDIT.

// Source: https://raw.github.com/awslabs/amazon-eks-ami/master/files/eni-max-pods.txt
var maxPodsPerNodeType = map[string]int{
	"c1.medium":    12,
	"c1.xlarge":    58,
	"c3.2xlarge":   58,
	"c3.4xlarge":   234,
	"c3.8xlarge":   234,
	"c3.large":     29,
	"c3.xlarge":    58,
	"c4.2xlarge":   58,
	"c4.4xlarge":   234,
	"c4.8xlarge":   234,
	"c4.large":     29,
	"c4.xlarge":    58,
	"c5.18xlarge":  737,
	"c5.2xlarge":   58,
	"c5.4xlarge":   234,
	"c5.9xlarge":   234,
	"c5.large":     29,
	"c5.xlarge":    58,
	"c5d.18xlarge": 737,
	"c5d.2xlarge":  58,
	"c5d.4xlarge":  234,
	"c5d.9xlarge":  234,
	"c5d.large":    29,
	"c5d.xlarge":   58,
	"cc2.8xlarge":  234,
	"cr1.8xlarge":  234,
	"d2.2xlarge":   58,
	"d2.4xlarge":   234,
	"d2.8xlarge":   234,
	"d2.xlarge":    58,
	"f1.16xlarge":  394,
	"f1.2xlarge":   58,
	"f1.4xlarge":   234,
	"g2.2xlarge":   58,
	"g2.8xlarge":   234,
	"g3.16xlarge":  737,
	"g3.4xlarge":   234,
	"g3.8xlarge":   234,
	"g3s.xlarge":   58,
	"h1.16xlarge":  737,
	"h1.2xlarge":   58,
	"h1.4xlarge":   234,
	"h1.8xlarge":   234,
	"hs1.8xlarge":  234,
	"i2.2xlarge":   58,
	"i2.4xlarge":   234,
	"i2.8xlarge":   234,
	"i2.xlarge":    58,
	"i3.16xlarge":  737,
	"i3.2xlarge":   58,
	"i3.4xlarge":   234,
	"i3.8xlarge":   234,
	"i3.large":     29,
	"i3.metal":     737,
	"i3.xlarge":    58,
	"m1.large":     29,
	"m1.medium":    12,
	"m1.small":     8,
	"m1.xlarge":    58,
	"m2.2xlarge":   118,
	"m2.4xlarge":   234,
	"m2.xlarge":    58,
	"m3.2xlarge":   118,
	"m3.large":     29,
	"m3.medium":    12,
	"m3.xlarge":    58,
	"m4.10xlarge":  234,
	"m4.16xlarge":  234,
	"m4.2xlarge":   58,
	"m4.4xlarge":   234,
	"m4.large":     20,
	"m4.xlarge":    58,
	"m5.12xlarge":  234,
	"m5.24xlarge":  737,
	"m5.2xlarge":   58,
	"m5.4xlarge":   234,
	"m5.large":     29,
	"m5.xlarge":    58,
	"m5a.12xlarge": 234,
	"m5a.24xlarge": 737,
	"m5a.2xlarge":  58,
	"m5a.4xlarge":  234,
	"m5a.large":    29,
	"m5a.xlarge":   58,
	"m5d.12xlarge": 234,
	"m5d.24xlarge": 737,
	"m5d.2xlarge":  58,
	"m5d.4xlarge":  234,
	"m5d.large":    29,
	"m5d.xlarge":   58,
	"p2.16xlarge":  234,
	"p2.8xlarge":   234,
	"p2.xlarge":    58,
	"p3.16xlarge":  234,
	"p3.2xlarge":   58,
	"p3.8xlarge":   234,
	"r3.2xlarge":   58,
	"r3.4xlarge":   234,
	"r3.8xlarge":   234,
	"r3.large":     29,
	"r3.xlarge":    58,
	"r4.16xlarge":  737,
	"r4.2xlarge":   58,
	"r4.4xlarge":   234,
	"r4.8xlarge":   234,
	"r4.large":     29,
	"r4.xlarge":    58,
	"r5.12xlarge":  234,
	"r5.24xlarge":  737,
	"r5.2xlarge":   58,
	"r5.4xlarge":   234,
	"r5.large":     29,
	"r5.xlarge":    58,
	"r5a.12xlarge": 234,
	"r5a.24xlarge": 737,
	"r5a.2xlarge":  58,
	"r5a.4xlarge":  234,
	"r5a.large":    29,
	"r5a.xlarge":   58,
	"r5d.12xlarge": 234,
	"r5d.24xlarge": 737,
	"r5d.2xlarge":  58,
	"r5d.4xlarge":  234,
	"r5d.large":    29,
	"r5d.xlarge":   58,
	"t1.micro":     4,
	"t2.2xlarge":   44,
	"t2.large":     35,
	"t2.medium":    17,
	"t2.micro":     4,
	"t2.nano":      4,
	"t2.small":     8,
	"t2.xlarge":    44,
	"t3.2xlarge":   44,
	"t3.large":     35,
	"t3.medium":    17,
	"t3.micro":     4,
	"t3.nano":      4,
	"t3.small":     8,
	"t3.xlarge":    44,
	"x1.16xlarge":  234,
	"x1.32xlarge":  234,
	"x1e.16xlarge": 234,
	"x1e.2xlarge":  58,
	"x1e.32xlarge": 234,
	"x1e.4xlarge":  58,
	"x1e.8xlarge":  58,
	"x1e.xlarge":   29,
	"z1d.12xlarge": 737,
	"z1d.2xlarge":  58,
	"z1d.3xlarge":  234,
	"z1d.6xlarge":  234,
	"z1d.large":    29,
	"z1d.xlarge":   58,
}
