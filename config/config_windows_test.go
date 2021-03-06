package config_test

import (
	"os"
	"path/filepath"

	"code.cloudfoundry.org/cfdev/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("config", func() {
	Describe("NewConfig", func() {
		Context("when CFDEV_HOME is not set", func() {
			var (
				oldHomeDrive string
				oldHomePath  string
			)

			BeforeEach(func() {
				oldHomeDrive = os.Getenv("HOMEDRIVE")
				oldHomePath = os.Getenv("HOMEPATH")
				os.Unsetenv("CFDEV_HOME")

				os.Setenv("HOMEDRIVE", "C:")
				os.Setenv("HOMEPATH", `\Users\some-home-dir`)
			})

			AfterEach(func() {
				os.Setenv("HOMEDRIVE", oldHomeDrive)
				os.Setenv("HOMEPATH", oldHomePath)
			})
			It("returns a config object with default values", func() {
				conf, err := config.NewConfig()
				Expect(err).NotTo(HaveOccurred())

				Expect(conf.BoshDirectorIP).To(Equal("10.144.0.4"))
				Expect(conf.CFRouterIP).To(Equal("10.144.0.34"))
				Expect(conf.HostIP).To(Equal("192.168.65.2"))
				Expect(conf.CFDevHome).To(Equal(filepath.Join(`C:\Users\some-home-dir`, ".cfdev")))
				Expect(conf.StateDir).To(Equal(filepath.Join(`C:\Users\some-home-dir`, ".cfdev", "state")))
				Expect(conf.StateBosh).To(Equal(filepath.Join(`C:\Users\some-home-dir`, ".cfdev", "state", "bosh")))
				Expect(conf.StateLinuxkit).To(Equal(filepath.Join(`C:\Users\some-home-dir`, ".cfdev", "state", "linuxkit")))
				Expect(conf.VpnKitStateDir).To(Equal(filepath.Join(`C:\Users\some-home-dir`, ".cfdev", "state", "vpnkit")))
				Expect(conf.CacheDir).To(Equal(filepath.Join(`C:\Users\some-home-dir`, ".cfdev", "cache")))
				Expect(conf.ServicesDir).To(Equal(filepath.Join(`C:\Users\some-home-dir`, ".cfdev", "services")))
				Expect(conf.LogDir).To(Equal(filepath.Join(`C:\Users\some-home-dir`, ".cfdev", "log")))
			})
		})
	})

	Context("when CFDEV_HOME is set", func() {
		BeforeEach(func() {
			os.Setenv("CFDEV_HOME", "some-cfdev-home")
		})

		AfterEach(func() {
			os.Unsetenv("CFDEV_HOME")
		})
		It("returns a config object with default values", func() {
			conf, err := config.NewConfig()
			Expect(err).NotTo(HaveOccurred())
			Expect(conf.BoshDirectorIP).To(Equal("10.144.0.4"))
			Expect(conf.CFRouterIP).To(Equal("10.144.0.34"))
			Expect(conf.HostIP).To(Equal("192.168.65.2"))
			Expect(conf.CFDevHome).To(Equal(filepath.Join("some-cfdev-home")))
			Expect(conf.StateDir).To(Equal(filepath.Join("some-cfdev-home", "state")))
			Expect(conf.StateBosh).To(Equal(filepath.Join("some-cfdev-home", "state", "bosh")))
			Expect(conf.StateLinuxkit).To(Equal(filepath.Join("some-cfdev-home", "state", "linuxkit")))
			Expect(conf.VpnKitStateDir).To(Equal(filepath.Join("some-cfdev-home", "state", "vpnkit")))
			Expect(conf.CacheDir).To(Equal(filepath.Join("some-cfdev-home", "cache")))
			Expect(conf.ServicesDir).To(Equal(filepath.Join("some-cfdev-home", "services")))
			Expect(conf.LogDir).To(Equal(filepath.Join("some-cfdev-home", "log")))
		})
	})
})
