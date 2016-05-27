module Dapp
  module CommonHelper
    def sha256(arg)
      Digest::SHA256.hexdigest Array(arg).compact.map(&:to_s).join(':::')
    end

    def kwargs(args)
      args.last.is_a?(Hash) ? args.pop : {}
    end
  end # CommonHelper
end # Dapp
