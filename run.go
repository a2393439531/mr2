// Copyright (c) 2019-present Cloud <cloud@txthinking.com>
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of version 3 of the GNU General Public
// License as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful, but
// WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
// General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package mr2

// RunServer .
func RunServer(address, password string, portPassword []string) error {
	s, err := NewServer(address, password, portPassword)
	if err != nil {
		return err
	}
	return s.ListenAndServe()
}

// RunClient .
func RunClient(server, password string, serverPort int64, serverDomain, clientServer string, tcpTimeout, tcpDeadline, udpDeadline int64) error {
	c := NewClient(server, password, serverPort, serverDomain, clientServer, tcpTimeout, tcpDeadline, udpDeadline)
	return c.Run()
}
