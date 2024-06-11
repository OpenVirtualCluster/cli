class Ovc < Formula
desc "CLI to install the virtual-cluster-operator Helm chart"
homepage "https://github.com/OpenVirtualCluster/cli"
version ""

if OS.mac?
  if Hardware::CPU.intel?
    url "https://github.com/OpenVirtualCluster/cli/releases/download//ovc-darwin-amd64"
    sha256 "a78374a7e512a93333c573b821a4b7fff14fbd53d68aa59ffa0343259a530250"
  elsif Hardware::CPU.arm?
    url "https://github.com/OpenVirtualCluster/cli/releases/download//ovc-darwin-arm64"
    sha256 "b9b175cc6f1483f6b4429b93abe5a1086ea529f44f9687b612f1dca18ce13d76"
  end
elsif OS.linux?
  if Hardware::CPU.intel?
    url "https://github.com/OpenVirtualCluster/cli/releases/download//ovc-linux-amd64"
    sha256 "3da9530c4c2d4185ea151e3cad532f76adc522cbff6d6854d8c0a5f126ccd344"
  elsif Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
    url "https://github.com/OpenVirtualCluster/cli/releases/download//ovc-linux-arm64"
    sha256 "a6a3b58e4f9dccab2758bb32ebee8c9e12362440607cfffb6d1d9c874bf4b904"
  end
end

def install
if OS.mac?
  if Hardware::CPU.intel?
    bin.install "ovc-darwin-amd64" => "ovc"
  elsif Hardware::CPU.arm?
    bin.install "ovc-darwin-arm64" => "ovc"
  end
elsif OS.linux?
  if Hardware::CPU.intel?
    bin.install "ovc-linux-amd64" => "ovc"
  elsif Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
    bin.install "ovc-linux-arm64" => "ovc"
  end
end
end

test do
  system "#{bin}/ovc", "version"
end
end
