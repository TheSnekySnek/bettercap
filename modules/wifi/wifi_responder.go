package wifi

import (
	"net"

	"github.com/bettercap/bettercap/packets"
)

func (mod *WiFiModule) sendProbeResponsePacket(client net.HardwareAddr, ssid string) {
	if err, pkt := packets.NewDot11ProbeResponse(client, mod.iface.HW, 1, ssid, mod.channel); err != nil {
		mod.Error("cloud not create probe response packet: %s", err)
	} else {
		mod.injectPacket(pkt)
	}
}

func (mod *WiFiModule) sendAuthenticationResponse(client net.HardwareAddr) {
	if err, pkt := packets.NewDot11AuthenticationResponse(client, mod.iface.HW, 1); err != nil {
		mod.Error("cloud not create authentication response packet: %s", err)
	} else {
		mod.injectPacket(pkt)
	}
}

func (mod *WiFiModule) sendAssociationResponse(client net.HardwareAddr) {
	if err, pkt := packets.NewDot11AssociationResponse(client, mod.iface.HW, 1); err != nil {
		mod.Error("cloud not create association response packet: %s", err)
	} else {
		mod.injectPacket(pkt)
	}
}

func (mod *WiFiModule) sendM1Wpa2(client net.HardwareAddr, seq uint16, replay uint64) {
	if err, pkt := packets.NewDot11M1Wpa2(client, mod.iface.HW, seq, replay); err != nil {
		mod.Error("cloud not create M1 packet: %s", err)
	} else {
		mod.injectPacket(pkt)
	}
}


func (mod *WiFiModule) startResponder() error {
	
	// if not already running, temporarily enable the pcap handle
	// for packet injection
	if !mod.Running() {
		if err := mod.Configure(); err != nil {
			return err
		}
		defer mod.handle.Close()
	}
	client, err := net.ParseMAC("c4:06:83:95:a9:ba")
	if err != nil {
		mod.Error("cloud not parse client mac: %s", err)
		return nil
	}
	mod.sendProbeResponsePacket(client, "TestAP")
	mod.sendProbeResponsePacket(client, "TestAP")
	mod.sendProbeResponsePacket(client, "TestAP")
	mod.sendAuthenticationResponse(client)
	mod.sendAuthenticationResponse(client)
	mod.sendAuthenticationResponse(client)
	mod.sendAssociationResponse(client)
	mod.sendAssociationResponse(client)
	mod.sendAssociationResponse(client)
	mod.sendM1Wpa2(client, uint16(1), uint64(1))
	mod.sendM1Wpa2(client, uint16(2), uint64(2))
	mod.sendM1Wpa2(client, uint16(3), uint64(3))

	return nil
}